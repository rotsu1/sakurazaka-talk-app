//
//  Models.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 8/1/2026.
//

import Foundation
import SwiftData

@Model
class Member {
    @Attribute(.unique) var id: String // ID from external DB
    var name: String
    var avatarUrl: String
    var joinOrder: Int
    var subscription: Subscription?
    
    // Helper to determine if they should be at the top
    var isSubscribed: Bool {
        subscription?.status == "active" ||
        Date() < (subscription?.expiryDate?.addingTimeInterval(60) ?? Date())
    }

    init(
        id: String,
        name: String,
        avatarUrl: String,
        joinOrder: Int,
        subscription: Subscription?
    ) {
        self.id = id
        self.name = name
        self.avatarUrl = avatarUrl
        self.joinOrder = joinOrder
        self.subscription = subscription
    }
}

@Model
class Subscription {
    var status: String // "active", "expired", etc.
    var expiryDate: Date?
    
    init(status: String, expiryDate: Date?) {
        self.status = status
        self.expiryDate = expiryDate
    }
}

@Model
class Blog {
    @Attribute(.unique) var id: String // ID from external DB
    var title: String
    var content: String
    var author: String
    var createdAt: Date

    init(
      id: String, 
      title: String, 
      content: String, 
      author: String, 
      createdAt: Date
    ) {
        self.id = id
        self.title = title
        self.content = content
        self.author = author
        self.createdAt = createdAt
    }
}

@Model
class OfficialNews {
    @Attribute(.unique) var id: String // ID from external DB
    var title: String
    var tag: String
    var content: String
    var createdAt: Date

    init(id: String, title: String, tag: String, content: String, createdAt: Date) {
        self.id = id
        self.title = title
        self.tag = tag
        self.content = content
        self.createdAt = createdAt
    }
}