//
//  NotificationService.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 8/1/2026.
//

import SwiftData
import Foundation

struct NotificationDTO: Codable {
    let id: Int
    let title: String
    let content: String
    let createdAt: Date
}

@MainActor
class NotificationService {
    let modelContext: ModelContext

    init(modelContext: ModelContext) {
        self.modelContext = modelContext
    }

    func syncNotifications() async throws {
        // 1. Fetch from Server
        guard let url = URL(string: "http://localhost:8080/notification/") else { return }
        let (data, _) = try await URLSession.shared.data(from: url)
        
        let decoder = JSONDecoder()
        decoder.keyDecodingStrategy = .convertFromSnakeCase
        decoder.dateDecodingStrategy = .iso8601
        let dtos = try decoder.decode([NotificationDTO].self, from: data)

        let allNotifications = try modelContext.fetch(FetchDescriptor<Notification>())
    
        var localMap = Dictionary(uniqueKeysWithValues: allNotifications.map { ($0.id, $0) })

        var unreadCount = try modelContext.fetch(FetchDescriptor<NotificationUnreadCount>()).first
        if unreadCount == nil {
            let newUnreadCount = NotificationUnreadCount(count: 0)
            modelContext.insert(newUnreadCount)
            unreadCount = newUnreadCount
        }

        // 2. Sync with SwiftData
        for dto in dtos {
            let stringID = String(dto.id)
            
            if let existingNotification = localMap[stringID] {
                // UPDATE existing if changed
                localMap.removeValue(forKey: stringID)

                if existingNotification.title != dto.title
                    || existingNotification.content != dto.content
                    || existingNotification.createdAt != dto.createdAt
                {
                    existingNotification.title = dto.title
                    existingNotification.content = dto.content
                    existingNotification.createdAt = dto.createdAt
                    existingNotification.isRead = false

                    unreadCount!.count += 1
                }
            } else {
                // INSERT new
                let newNotification = Notification(
                    id: stringID,
                    title: dto.title,
                    content: dto.content,
                    createdAt: dto.createdAt,
                    isRead: false
                )
                modelContext.insert(newNotification)
                unreadCount!.count += 1
            }
        }

        // 3. Process Deletions (The "Remaining" Items)
        // Anything still left in `localMap` was NOT in the server response.
        for (_, notificationToDelete) in localMap {
            // If we are deleting an unread notification, we should decrease the badge count
            if !notificationToDelete.isRead && unreadCount != nil {
                unreadCount!.count -= 1
            }
            
            modelContext.delete(notificationToDelete)
        }
        
        // 4. Save changes
        try modelContext.save()
    }
}
