//
//  BlogService.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 8/1/2026.
//

import SwiftData
import Foundation

struct BlogDTO: Codable {
    let id: Int
    let title: String
    let content: String
    let author: String
    let createdAt: Date
}

@MainActor
class BlogService {
    let modelContext: ModelContext

    init(modelContext: ModelContext) {
        self.modelContext = modelContext
    }

    func syncBlogs() async throws {
        // 1. Fetch from Server
        var components = URLComponents(string: "http://localhost:8080/blog/")!
        components.queryItems = [
            URLQueryItem(name: "status", value: "verified")
        ]
        guard let url = components.url else { return }

        let (data, _) = try await URLSession.shared.data(from: url)
        
        let decoder = JSONDecoder()
        decoder.keyDecodingStrategy = .convertFromSnakeCase
        decoder.dateDecodingStrategy = .iso8601
        let dtos = try decoder.decode([BlogDTO].self, from: data)

        // 2. Sync with SwiftData
        for dto in dtos {
            let stringID = String(dto.id)
            
            // Check if member exists
            let fetchDescriptor = FetchDescriptor<Blog>(
                predicate: #Predicate { $0.id == stringID }
            )
            
            if let existingBlog = try modelContext.fetch(fetchDescriptor).first {
                // UPDATE existing if changed
                if existingBlog.title != dto.title
                    || existingBlog.content != dto.content
                    || existingBlog.author != dto.author
                    || existingBlog.createdAt != dto.createdAt
                {
                    existingBlog.title = dto.title
                    existingBlog.content = dto.content
                    existingBlog.author = dto.author
                    existingBlog.createdAt = dto.createdAt
                }
            } else {
                // INSERT new
                let newBlog = Blog(
                    id: stringID,
                    title: dto.title,
                    content: dto.content,
                    author: dto.author,
                    createdAt: dto.createdAt
                )
                modelContext.insert(newBlog)
            }
        }
        
        // 3. Save changes
        try modelContext.save()
    }
}
