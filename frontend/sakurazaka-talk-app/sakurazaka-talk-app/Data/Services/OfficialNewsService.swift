//
//  OfficialNewsService.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 8/1/2026.
//

import SwiftData
import Foundation

struct OfficialNewsDTO: Codable {
    let id: Int
    let title: String
    let tag: String
    let content: String
    let createdAt: Date
}

@MainActor
class OfficialNewsService {
    let modelContext: ModelContext

    init(modelContext: ModelContext) {
        self.modelContext = modelContext
    }

    func syncOfficialNews() async throws {
        // 1. Fetch from Server
        guard let url = URL(string: "http://localhost:8080/official_news/") else { return }
        let (data, _) = try await URLSession.shared.data(from: url)
        
        let decoder = JSONDecoder()
        decoder.keyDecodingStrategy = .convertFromSnakeCase
        decoder.dateDecodingStrategy = .iso8601
        let dtos = try decoder.decode([OfficialNewsDTO].self, from: data)

        let allOfficialNews = try modelContext.fetch(FetchDescriptor<OfficialNews>())

        var localMap = Dictionary(uniqueKeysWithValues: allOfficialNews.map { ($0.id, $0) })

        // 2. Sync with SwiftData
        for dto in dtos {
            let stringID = String(dto.id)
            
            if let existingOfficialNews = localMap[stringID] {
                // UPDATE existing if changed
                localMap.removeValue(forKey: stringID)
                
                if existingOfficialNews.title != dto.title
                    || existingOfficialNews.tag != dto.tag
                    || existingOfficialNews.content != dto.content
                    || existingOfficialNews.createdAt != dto.createdAt
                {
                    existingOfficialNews.title = dto.title
                    existingOfficialNews.tag = dto.tag
                    existingOfficialNews.content = dto.content
                    existingOfficialNews.createdAt = dto.createdAt
                }
            } else {
                // INSERT new
                let newOfficialNews = OfficialNews(
                    id: stringID,
                    title: dto.title,
                    tag: dto.tag,
                    content: dto.content,
                    createdAt: dto.createdAt
                )
                modelContext.insert(newOfficialNews)
            }
        }

        // 3. Process Deletions (The "Remaining" Items)
        for (_, officialNewsToDelete) in localMap {
            modelContext.delete(officialNewsToDelete)
        }
        
        // 4. Save changes
        try modelContext.save()
    }
}
