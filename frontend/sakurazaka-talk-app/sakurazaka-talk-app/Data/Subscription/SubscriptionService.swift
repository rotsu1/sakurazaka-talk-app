//
//  SubscriptionService.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 9/1/2026.
//

import StoreKit
import SwiftData
import Foundation
import Combine

@MainActor
class SubscriptionService: ObservableObject {
    @Published var subscriptions: [Subscription] = []

    // Simulate purchasing a subscription for a member
    func subscribe(to member: Member, context: ModelContext) async throws {
        
        // For now, simulate successful subscription (1 month)
        let expiryDate = Calendar.current.date(byAdding: .month, value: 1, to: Date())
        
        if let existingSubscription = member.subscription {
            existingSubscription.status = "active"
            existingSubscription.expiryDate = expiryDate
        } else {
            let newSubscription = Subscription(status: "active", expiryDate: expiryDate)
            member.subscription = newSubscription
        }

        try? context.save()
        
        // Send subscription detail to backend
        do {
            try await syncSubscriptionToBackend(member: member, expiryDate: expiryDate!)
        } catch {
            print("Failed to sync subscription to backend: \(error)")
        }
    }
    
    // Sync subscription to backend
    private func syncSubscriptionToBackend(member: Member, expiryDate: Date) async throws {
        let urlString = "http://localhost:8080/talk_user_member/" // Replace with actual API URL
        guard let url = URL(string: urlString) else { return }
        
        // Assuming current user ID is available or handled by auth token
        let userID = 1 // TODO: Get actual user ID
        
        // Ensure Member ID is convertible to Int, or handle error if string
        let memberID = Int(member.id) ?? 0 
        
        let payload: [String: Any] = [
            "user_id": userID,
            "member_id": memberID,
            "status": "active",
            "expires_at": ISO8601DateFormatter().string(from: expiryDate)
        ]
        
        var request = URLRequest(url: url)
        request.httpMethod = "POST"
        request.setValue("application/json", forHTTPHeaderField: "Content-Type")
        request.httpBody = try JSONSerialization.data(withJSONObject: payload)
        
        let (data, response) = try await URLSession.shared.data(for: request)
    
    guard let httpResponse = response as? HTTPURLResponse else {
        throw URLError(.badServerResponse)
    }

    if !(200...299).contains(httpResponse.statusCode) {
        // THIS IS KEY: Print what the server actually says
        if let serverErrorMessage = String(data: data, encoding: .utf8) {
            print("❌ Server Error (\(httpResponse.statusCode)): \(serverErrorMessage)")
        }
        throw URLError(.badServerResponse)
    }
    }
    
    func isSubscribed(to member: Member) -> Bool {
        guard let subscription = member.subscription, 
              let expiry = subscription.expiryDate else {
            return false
        }
        return subscription.status == "active" && expiry > Date()
    }
}
