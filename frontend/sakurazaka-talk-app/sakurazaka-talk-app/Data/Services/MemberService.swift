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

        // 2. Sync with SwiftData
        for dto in dtos {
            let stringID = String(dto.id)
            
            // Check if member exists
            let fetchDescriptor = FetchDescriptor<Member>(
                predicate: #Predicate { $0.id == stringID }
            )
            
            if let existingMember = try modelContext.fetch(fetchDescriptor).first {
                // UPDATE existing if changed
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
        
        // 3. Save changes
        try modelContext.save()
    }
}
