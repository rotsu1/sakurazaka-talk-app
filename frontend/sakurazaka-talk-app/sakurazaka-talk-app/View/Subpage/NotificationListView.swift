//
//  NotificationListView.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 3/1/2026.
//

import SwiftUI
import SwiftData

struct NotificationItem: Identifiable, Hashable {
    let id = UUID()
    let title: String
    let content: String
    let timestamp: Date
}

let calendar = Calendar.current

struct NotificationListView: View {
    @Query(sort: \Notification.createdAt) private var notifications: [Notification]
    @Query private var unreadCounts: [NotificationUnreadCount]

    @Environment(\.modelContext) private var modelContext
    
    let calendar = Calendar.current

    var body: some View {
        VStack {
            HeaderView(title: "お知らせ", icons: false, isBlog: false, isSubpage: true)
            
            ScrollView {
                LazyVStack(alignment: .leading, spacing: 12) {
                    ForEach(notifications) { notification in
                        NavigationLink(
                            destination: NotificationView(notification: notification)
                        ) {
                            let dateString = formatterSimple.string(
                                from: notification.createdAt
                            )
                            if notification.isRead {
                                notificationContent(
                                    date: dateString, 
                                    title: notification.title, 
                                    content: notification.content,
                                )
                            } else {
                                ZStack(alignment: .topTrailing) {
                                    notificationContent(
                                        date: dateString, 
                                        title: notification.title, 
                                        content: notification.content,
                                    )
                                    Circle().fill(sakuraPink).frame(width: 15, height: 15)
                                        .offset(x: 5, y: -5)
                                }
                            }
                        }
                        .simultaneousGesture(TapGesture().onEnded {
                            if !notification.isRead {
                                if let unreadCount = unreadCounts.first {
                                    unreadCount.count -= 1
                                }
                            }
                            notification.isRead = true
                            try? modelContext.save()
                        })
                    }
                }
                .padding()
            }
            .navigationBarHidden(true) 
            .navigationBarBackButtonHidden(true)
        }
        .task {
            do {
                try await NotificationService(modelContext: modelContext).syncNotifications()
            } catch {
                print("Error syncing notifications: \(error)")
            }
        }
    }

    func notificationContent(date: String, title: String, content: String) -> some View {
        VStack(alignment: .leading, spacing: 8) {
            Text(date)
                .foregroundColor(Color(white: 0.6))
                .font(.system(size: 13, weight: .regular))
                .lineLimit(1)
            Text(title)
                .foregroundColor(Color(white: 0.5))
                .font(.system(size: 17, weight: .regular))
                .lineLimit(1)
            Text(content)
                .foregroundColor(Color(white: 0.3))
                .font(.system(size: 14, weight: .medium))
                .lineLimit(1)
        }
        .padding()
        .frame(maxWidth: .infinity, alignment: .leading)
        .background(Color.rgb(red: 247, green: 247, blue: 247))
    }
}

#Preview {
    NotificationListView()
}