//
//  Header.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 3/1/2026.
//

import SwiftUI

struct HeaderView: View {
    let title: String
    let icons: Bool
    let isBlog: Bool
    let isSubpage: Bool
    @Environment(\.dismiss) var dismiss

    var body: some View {
        ZStack {
            Text(title)
                .font(.headline)

            HStack {
                if isSubpage {
                    Button(action: {
                        dismiss()
                    }) {
                        Image(systemName: "chevron.left")
                            .font(.system(size: 24))
                            .padding(.leading, 8)
                    }
                }

                Spacer()

                if icons {
                    HStack(spacing: 16) {
                        if isBlog {
                            Image(systemName: "arrow.up.arrow.down")
                                .font(.system(size: 24))
                        }
                        NavigationLink(destination: NotificationView()) {
                            ZStack(alignment: .topTrailing){
                                Image(systemName: "bell")
                                    .font(.system(size: 24))

                                Text("20")
                                    .font(.caption)
                                    .foregroundColor(Color.white)
                                    .frame(width: 20, height: 20) 
                                    .background(Circle().fill(sakuraPink))
                                    .padding(4)
                                    .offset(x: 12, y: -12)
                            }
                        }
                        NavigationLink(destination: SettingsView()) {
                            Image(systemName: "gearshape")
                                .font(.system(size: 24))
                        }
                    }
                }
            }
            .padding(.horizontal, 8)
        }
        .padding(.bottom, 16)
        .foregroundColor(sakuraPink)
        .overlay(
            Rectangle()
                .fill(
                    LinearGradient(
                        colors: [
                            Color.white,
                            Color.rgb(red: 213, green: 0, blue: 255)
                        ],
                        startPoint: .leading,
                        endPoint: .trailing
                    )
                )
                .frame(height: 3),
            alignment: .bottom
        )
    }
}

#Preview {
    HeaderView(title: "トーク", icons: true, isBlog: false, isSubpage: false)
}
