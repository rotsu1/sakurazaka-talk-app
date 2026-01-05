//
//  NotificationView.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 4/1/2026.
//

import SwiftUI

struct NotificationView: View {
    let notificationItem: NotificationItem

    var body: some View {
        VStack {
            HeaderView(title: "お知らせ", icons: false, isBlog: false, isSubpage: true)

            ScrollView {
                VStack(alignment: .leading, spacing: 16) {
                    Text(
                        formatterDetailed.string(
                            from: notificationItem.timestamp)
                        )
                        .foregroundColor(Color(white: 0.6))
                        .font(.system(size: 13, weight: .regular))
                        .frame(maxWidth: .infinity, alignment: .trailing)
                    Text(notificationItem.title)
                        .font(.system(size: 18, weight: .medium))
                        .frame(maxWidth: .infinity, alignment: .leading)
                    Text(notificationItem.content)
                        .font(.system(size: 16, weight: .regular))
                }
                .padding()
                .scrollIndicators(.hidden)
            }
            .navigationBarHidden(true)
            .navigationBarBackButtonHidden(true)
        }
    }
}
