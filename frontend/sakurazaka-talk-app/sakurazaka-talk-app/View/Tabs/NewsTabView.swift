//
//  NewsTabView.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 4/1/2026.
//

import SwiftUI
import SwiftData

enum NewsTabs {
    case first
    case second
}

struct NewsItem: Identifiable {
    let id = UUID()
    let title: String
    let tag: String
    let content: String
    let createdAt: Date
}

let fanclubNews = [
    NewsItem(
        title: "「さくみみ」#536公開！",
        tag: "ファンクラブ",
        content: "",
        createdAt: calendar.date(from: DateComponents(year: 2026, month: 1, day: 2))!,
    ),
    NewsItem(
        title: "「MANAGER'S DIARY」更新！",
        tag: "ファンクラブ",
        content: "",
        createdAt: calendar.date(from: DateComponents(year: 2026, month: 1, day: 1))!,
    ),
    NewsItem(
        title: "「2026年 お正月SPECIAL SITE」にて年賀状と書き初めを公開！",
        tag: "ファンクラブ",
        content: "",
        createdAt: calendar.date(from: DateComponents(year: 2026, month: 1, day: 1))!,
    ),
    NewsItem(
        title: "「MANAGER'S DIARY」更新！",
        tag: "ファンクラブ",
        content: "",
        createdAt: calendar.date(from: DateComponents(year: 2025, month: 12, day: 31))!,
    ),
    NewsItem(
        title: "「さくみみ」#535公開！",
        tag: "ファンクラブ",
        content: "",
        createdAt: calendar.date(from: DateComponents(year: 2025, month: 12, day: 30))!,
    ),
]

struct NewsTabView: View {
    @Query(sort: \OfficialNews.createdAt, order: .reverse) private var officialNews: [OfficialNews]
    @Environment(\.modelContext) private var modelContext

    @State private var selectedTab: NewsTabs = .first

    var body: some View {
        VStack {
            HeaderView(title: "お知らせ", icons: true, isBlog: false, isSubpage: false)
            NewsTabButtons(selectedTab: $selectedTab)

            ScrollView {
                VStack(spacing: 8) {
                    if selectedTab == .first {
                            ForEach(officialNews) { news in
                            NewsItemView(
                                title: news.title,
                                tag: news.tag,
                                content: news.content,
                                createdAt: news.createdAt,
                            )
                        }
                    } else {
                        ForEach(fanclubNews) { news in
                            NewsItemView(
                                title: news.title,
                                tag: news.tag,
                                content: news.content,
                                createdAt: news.createdAt,
                            )
                        }
                    }
                }
            }
            .scrollIndicators(.hidden)
            .navigationBarHidden(true) 
            .navigationBarBackButtonHidden(true)
            .padding()

            Spacer().frame(height: 100)
        }
        .task {
            do {
                try await OfficialNewsService(modelContext: modelContext).syncOfficialNews()
            } catch {
                print("Error syncing official news: \(error)")
            }
        }
    }
}

struct NewsTabButtons: View {

    @Binding var selectedTab: NewsTabs

    var body: some View {
        VStack(spacing: 0) {
            GeometryReader { geo in
                ZStack(alignment: .bottomLeading) {
                    HStack(spacing: 0) {
                        tabButton(title: "櫻坂46 OFFICIAL", tab: .first)
                        tabButton(title: "櫻坂46 FANCLUB", tab: .second)
                    }

                    // Bottom indicator
                    Rectangle()
                        .fill(sakuraPink)
                        .frame(width: geo.size.width / 2, height: 3)
                        .offset(x: selectedTab == .first ? 0 : geo.size.width / 2, y: 0)
                        .animation(.easeInOut(duration: 0.25), value: selectedTab)
                }
            }
            .frame(height: 44)
        }
    }

    private func tabButton(title: String, tab: NewsTabs) -> some View {
        let isSelected = selectedTab == tab

        return Button {
            selectedTab = tab
        } label: {
            ZStack {
                // 1. Regular weight (always there, just fades)
            Text(title)
                .font(.system(size: 16, weight: .regular))
                .foregroundColor(sakuraPink)
                .opacity(isSelected ? 0 : 1)
            
            // 2. Semibold weight (already loaded in memory)
            Text(title)
                .font(.system(size: 16, weight: .semibold))
                .foregroundColor(sakuraPink)
                .opacity(isSelected ? 1 : 0)

            }
            .frame(maxWidth: .infinity, maxHeight: .infinity)
            .contentShape(Rectangle())
        }
    }
}

struct NewsItemView: View {
    let title: String
    let tag: String
    let content: String
    let createdAt: Date

    var body: some View {
        NavigationLink(destination: NewsView(content: content)) {
            VStack(spacing: 4) {
                HStack {
                    Text(tag)
                        .foregroundColor(tagTextColor[tag] ?? Color.white)
                        .font(.system(size: 12, weight: .bold))
                        .padding(4)
                        .frame(width: 80)
                        .background(tagColor[tag] ?? Color.black)
                        .cornerRadius(4)

                    Spacer()

                    Text(formatterSimple.string(from: createdAt))
                        .font(.system(size: 10, weight: .regular))
                        .foregroundColor(Color(white: 0.6))
                }

                Text(title)
                    .font(.system(size: 16, weight: .regular))
                    .foregroundColor(Color(white: 0.5))
                    .frame(maxWidth: .infinity, alignment: .leading)
                    .multilineTextAlignment(.leading)
                    .lineLimit(2)
            }
            .padding()
            .frame(maxWidth: .infinity, alignment: .leading)
            .background(Color.rgb(red: 247, green: 247, blue: 247))
        }
    }
}

#Preview {
    NewsTabView()
}
