//
//  MemberService.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 8/1/2026.
//

import SwiftData
import Foundation

struct MemberDTO: Codable {
    let id: Int
    let name: String
    let avatarUrl: String
    let generation: Int
}

@MainActor
class MemberService {
    let modelContext: ModelContext

    init(modelContext: ModelContext) {
        self.modelContext = modelContext
    }

    func syncMembers() async throws {
        // 1. Fetch from Server
        guard let url = URL(string: "http://localhost:8080/member/") else { return }
        let (data, _) = try await URLSession.shared.data(from: url)
        
        let decoder = JSONDecoder()
        decoder.keyDecodingStrategy = .convertFromSnakeCase
        let dtos = try decoder.decode([MemberDTO].self, from: data)

        let allMembers = try modelContext.fetch(FetchDescriptor<Member>())
        var localMap = Dictionary(uniqueKeysWithValues: allMembers.map { ($0.id, $0) })

        // 2. Sync with SwiftData
        for dto in dtos {
            let stringID = String(dto.id)
            
            if let existingMember = localMap[stringID] {
                // UPDATE existing if changed
                localMap.removeValue(forKey: stringID)
                
                if existingMember.name != dto.name
                    || existingMember.avatarUrl != dto.avatarUrl
                    || existingMember.joinOrder != dto.generation
                {
                    existingMember.name = dto.name
                    existingMember.avatarUrl = dto.avatarUrl
                    existingMember.joinOrder = dto.generation
                }
            } else {
                // INSERT new
                let newMember = Member(
                    id: stringID,
                    name: dto.name,
                    avatarUrl: dto.avatarUrl,
                    joinOrder: dto.generation,
                    subscription: nil
                )
                modelContext.insert(newMember)
            }
        }

        // 3. Process Deletions (The "Remaining" Items)
        for (_, memberToDelete) in localMap {
            modelContext.delete(memberToDelete)
        }
        
        // 4. Save changes
        try modelContext.save()
    }
}
