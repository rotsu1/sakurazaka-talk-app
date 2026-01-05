//
//  NewsTabView.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 4/1/2026.
//

import SwiftUI

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

let officialNews = [
    NewsItem(
        title: "1月4日(日)17:00～TBS「バナナマンのせっかくグルメ！！」に松田里奈が出演！",
        tag: "メディア",
        content: "",
        createdAt: calendar.date(from: DateComponents(year: 2026, month: 1, day: 4))!,
    ),
    NewsItem(
        title: "「櫻坂チャンネル」にて「【裏側】シンガポールで開催の「AFASG25」に櫻坂46が出演！オフの時間も満喫！【Vlog】」を公開！",
        tag: "メディア",
        content: "",
        createdAt: calendar.date(from: DateComponents(year: 2025, month: 12, day: 29))!,
    ),
    NewsItem(
        title: "今年リリースされた楽曲を振り返るStationheadリスニングパーティー開催決定！",
        tag: "リリース",
        content: "",
        createdAt: calendar.date(from: DateComponents(year: 2025, month: 12, day: 27))!,
    ),
    NewsItem(
        title: "LAWSON 50th Anniversary presents Special LIVE 〜櫻坂46 / 日向坂46 〜の配信が決定！配信が決定！配信視聴チケットは、12月26日（金）15：00より販売スタート！",
        tag: "イベント情報",
        content: "",
        createdAt: calendar.date(from: DateComponents(year: 2025, month: 12, day: 26))!,
    ),
    NewsItem(
        title: "「TopYellNEO2025~2026」（12月27日(土)発売）の表紙・巻頭に石森璃花が、中面に大沼晶保と増本キラが登場！",
        tag: "メディア",
        content: "",
        createdAt: calendar.date(from: DateComponents(year: 2025, month: 12, day: 26))!,
    ),
    NewsItem( 
        title: "「20±SWEET2026 JANUARY」（12月26日(金)発売）に遠藤理子、谷口愛季、稲熊ひな、山川宇彩が登場！",
        tag: "メディア",
        content: "",
        createdAt: calendar.date(from: DateComponents(year: 2025, month: 12, day: 25))!,
    ),
    NewsItem(
        title: "アニプレックスYouTubeチャンネルにて公開の「応援大使（谷口愛季）のお気に入りアニメ『...",
        tag: "メディア",
        content: "",
        createdAt: calendar.date(from: DateComponents(year: 2025, month: 12, day: 20))!,
    ),
    NewsItem(
        title: "2026年放送のテレビ朝日「あざとくて何が悪いの？」内「あざと連ドラ」に田村保乃の出演が...",
        tag: "メディア",
        content: "",
        createdAt: calendar.date(from: DateComponents(year: 2025, month: 12, day: 19))!,
    ),
    NewsItem(
        title: "12月26日(金)20:00～FOD／Prime Videoにて配信開始のオリジナルドラマ『にこたま』に田村保乃...",
        tag: "メディア",
        content: "",
        createdAt: calendar.date(from: DateComponents(year: 2025, month: 12, day: 19))!,
    ),
    NewsItem(
        title: "「櫻坂チャンネル」にて「【#二期生ずっと一緒】二期生考案 MV風『紋白蝶が確かに飛んでいた』」を公開！",
        tag: "メディア",
        content: "",
        createdAt: calendar.date(from: DateComponents(year: 2025, month: 12, day: 17))!,
    ),
    NewsItem(
        title: "櫻坂46「SAKURAZAKA46 SPORTS FESTIVAL supported by AEON CARD」、イオンカード(櫻坂46)がイベントを支援！",
        tag: "イベント情報",
        content: "",
        createdAt: calendar.date(from: DateComponents(year: 2025, month: 12, day: 17))!,
    ),
]

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
