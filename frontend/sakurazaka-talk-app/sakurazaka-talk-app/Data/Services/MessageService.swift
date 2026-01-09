//
//  MessageService.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 9/1/2026.
//

import SwiftData
import Foundation

struct MessageDTO: Codable {
    let id: Int
    let memberId: Int
    let type: String
    let content: String
    let createdAt: Date
}

@MainActor
class MessageService {
    let modelContext: ModelContext

    init(modelContext: ModelContext) {
        self.modelContext = modelContext
    }

    func syncMessages() async throws {
        // 1. Fetch from Server
        guard let url = URL(string: "http://localhost:8080/message/") else { return }
        let (data, _) = try await URLSession.shared.data(from: url)
        
        let decoder = JSONDecoder()
        decoder.keyDecodingStrategy = .convertFromSnakeCase
        decoder.dateDecodingStrategy = .iso8601
        let dtos = try decoder.decode([MessageDTO].self, from: data)

        // 2. Prepare Local Data
        let allMessages = try modelContext.fetch(FetchDescriptor<Message>())
        let allMembers = try modelContext.fetch(FetchDescriptor<Member>())
        
        var localMessageMap = Dictionary(uniqueKeysWithValues: allMessages.map { ($0.id, $0) })
        // Map members by ID for quick lookup
        let memberMap = Dictionary(uniqueKeysWithValues: allMembers.map { ($0.id, $0) })

        // 3. Sync Logic
        for dto in dtos {
            let stringID = String(dto.id)
            let memberID = String(dto.memberId)
            
            // Critical: Ensure the member exists in local DB before creating a message
            guard let member = memberMap[memberID] else {
                print("Skipping message \(stringID): Member \(memberID) not found locally.")
                continue
            }

            if let existingMessage = localMessageMap[stringID] {
                localMessageMap.removeValue(forKey: stringID)
                
                // Update properties if changed
                if existingMessage.content != dto.content {
                    existingMessage.content = dto.content
                    // Re-download data if the URL changed
                    if dto.type != "text" {
                        await saveMessageData(message: existingMessage, from: URL(string: dto.content)!)
                    }
                }
                existingMessage.type = dto.type
                existingMessage.createdAt = dto.createdAt
                existingMessage.member = member
                
            } else {
                // INSERT new message
                let newMessage = Message(
                    id: stringID,
                    member: member,
                    type: dto.type,
                    content: dto.content,
                    createdAt: dto.createdAt
                )
                
                // Download media if necessary
                if dto.type != "text", let mediaURL = URL(string: dto.content) {
                    await saveMessageData(message: newMessage, from: mediaURL)
                }
                
                modelContext.insert(newMessage)
            }
        }
        
        // 4. Deletions
        for (_, messageToDelete) in localMessageMap {
            modelContext.delete(messageToDelete)
        }
        
        try modelContext.save()
    }

    func saveMessageData(message: Message, from url: URL) async {
        do {
            let (data, _) = try await URLSession.shared.data(from: url)
            message.data = data
        } catch {
            print("Failed to download media for message \(message.id): \(error)")
        }
    }
}
