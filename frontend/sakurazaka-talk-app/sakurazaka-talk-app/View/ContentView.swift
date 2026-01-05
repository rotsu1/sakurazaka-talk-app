//
//  ContentView.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 3/1/2026.
//

import SwiftUI

enum Tab: String, CaseIterable {
    case talk = "Talk"
    case blog = "Blog"
    case news = "News"
    case official = "Official"
    case fanclub = "Fanclub"

    var icon: String {
        switch self {
        case .talk: return "bubble.right"
        case .blog: return "text.page"
        case .news: return "megaphone"
        case .official: return "triangleshape"
        case .fanclub: return "leaf" // Representing the seed/flower logo
        }
    }
}

struct CustomTabBar: View {
    @Binding var selectedTab: Tab
    @Environment(\.openURL) var openURL

    let activeColor = Color(red: 0.85, green: 0.45, blue: 0.55) // Sakura Pink
    let inactiveColor = Color.gray.opacity(0.6)

    var body: some View {
        VStack(spacing: 0) {
            Divider()
            
            HStack(spacing: 0) {
                ForEach(Tab.allCases, id: \.self) { tab in
                    Button {
                        if tab == .official {
                            // Open external browser
                            if let url = URL(string: "https://sakurazaka46.com/") {
                                openURL(url)
                            }
                        } else {
                            // Normal tab switching
                            selectedTab = tab
                        }
                    } label: {
                        VStack(spacing: 4) {
                            Image(systemName: tab.icon)
                                .font(.system(size: 22, weight: .light))
                            
                            Text(tab.rawValue)
                                .font(.system(size: 10, weight: .medium))
                        }
                        .frame(maxWidth: .infinity)
                        .padding(.vertical, 8)
                        // Note: Official tab won't ever look "active" since it redirects away
                        .foregroundStyle(selectedTab == tab ? activeColor : inactiveColor)
                    }
                }
            }
            .padding(.bottom, 34) // Standard iPhone bottom safe area height
            .background(Color.white)
        }
    }
}


struct ContentView: View {
    @State private var selectedTab: Tab = .talk

    var body: some View {
        ZStack(alignment: .bottom) {
            Group {
                switch selectedTab {
                case .talk: TalkTabView()
                case .blog: BlogTabView()
                case .news: NewsTabView()
                case .official: EmptyView()
                case .fanclub: FanclubTabView()
                }
            }
            .frame(maxWidth: .infinity, maxHeight: .infinity)
            
            CustomTabBar(selectedTab: $selectedTab)
        }
        .ignoresSafeArea(edges: .bottom)
    }
}

#Preview {
    ContentView()
}
