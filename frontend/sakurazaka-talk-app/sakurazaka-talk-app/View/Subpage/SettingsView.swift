//
//  SettingsView.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 3/1/2026.
//

import SwiftUI

struct SettingsView: View {
    var body: some View {
        // Use a VStack with spacing 0 to ensure the header and content touch
        VStack(spacing: 0) {
            HeaderView(title: "設定", icons: false, isBlog: false, isSubpage: true)
            
            // This ScrollView will contain the gray background
            ScrollView {
                VStack(spacing: 0) {
                    Group {
                        SettingsItem(
                            title: "アカウント", 
                            icon: "person", 
                            description: "プロフィール編集、ログインアカウント設定", 
                            link: EmptyView()
                        )
                        Rectangle()
                            .fill(Color.gray.opacity(0.2))
                            .frame(height: 1)
                            .padding(.horizontal, 20)
                        SettingsItem(
                            title: "通知", 
                            icon: "info.circle", 
                            description: "トーク、お知らせなどの通知設定", 
                            link: EmptyView()
                        )
                        Rectangle()
                            .fill(Color.gray.opacity(0.2))
                            .frame(height: 1)
                            .padding(.horizontal, 20)
                        SettingsItem(
                            title: "トーク", 
                            icon: "bubble", 
                            description: "トークに関する設定", 
                            link: EmptyView()
                        )
                        Rectangle()
                            .fill(Color.gray.opacity(0.2))
                            .frame(height: 1)
                            .padding(.horizontal, 20)
                        SettingsItem(
                            title: "定期購入", 
                            icon: "checkmark.seal.text.page", 
                            description: "定期購入の確認", 
                            link: EmptyView()
                        )
                        Rectangle()
                            .fill(Color.gray.opacity(0.2))
                            .frame(height: 1)
                            .padding(.horizontal, 20)
                        SettingsItem(
                            title: "サポート", 
                            icon: "person", 
                            description: "アプリの使い方、お問い合わせ", 
                            link: EmptyView()
                        )
                    }
                    .background(Color.white) // Makes the items white
                    
                    SettingsItem(
                        title: "「櫻坂46ファンクラブ」連携", 
                        icon: "person.3", 
                        description: "", 
                        link: EmptyView()
                    )
                    .background(Color.white) // Separate white block
                    .padding(.top, 16)
                }
                HStack {
                    Text("アプリバージョン")
                    Spacer()
                    Text("1.0.0")
                }
                .font(.system(size: 14, weight: .medium))
                .foregroundColor(Color(white: 0.4))
                .padding()
            }
            .background(Color(white: 0.97)) // This is the light gray background for the page
        }
        .navigationBarHidden(true)
        .navigationBarBackButtonHidden(true)
    }
}

struct SettingsItem<Destination: View>: View {
    let title: String
    let icon: String
    let description: String
    let link: Destination

    var body: some View {
        NavigationLink(destination: link) {
            HStack {
                HStack(spacing: 16) {
                    Image(systemName: icon)
                        .font(.system(size: 24))
                        .foregroundColor(sakuraPink)
                        .frame(width: 32)
                    VStack(alignment: .leading, spacing: 4) {
                        Text(title)
                            .font(.system(size: 16, weight: .medium))
                            .foregroundColor(Color(white: 0.5))
                        if description != "" {
                            Text(description)
                                .font(.system(size: 14, weight: .regular))
                        }
                    }
                }

                Spacer()

                Image(systemName: "chevron.right")
                    .font(.system(size: 20))
                    .foregroundColor(sakuraPink)
            }
            .frame(minHeight: 30)
            .padding(.horizontal, 16)
            .padding(.vertical, 10)
        }
    }
}

#Preview {
    SettingsView()
}