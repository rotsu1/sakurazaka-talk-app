//
//  Header.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 3/1/2026.
//

import SwiftUI

struct Header: View {
    let title: String

    var body: some View {
        ZStack {
            Text(title)
                .font(.headline)

            HStack {
                Spacer()

                HStack {
                    NavigationLink(destination: NotificationView()) {
                        ZStack(alignment: .topTrailing){
                            Image(systemName: "bell")
                                .font(.system(size: 24))

                            Text("20")
                                .font(.caption)
                                .foregroundColor(Color.white)
                                .frame(width: 20, height: 20) 
                                .background(Circle().fill(Color.rgb(red: 255, green: 149, blue: 228)))
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
            .padding(.horizontal, 8)
        }
        .padding(.bottom, 16)
        .foregroundColor(Color.rgb(red: 255, green: 149, blue: 228))
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
    Header(title: "トーク")
}
