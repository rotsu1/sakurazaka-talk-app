//
//  TalkTabView.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 3/1/2026.
//

import SwiftUI

struct MemberGroup: Identifiable {
    let id = UUID()
    let generation: String
    let names: [String]
}

let memberData: [MemberGroup] = [
    MemberGroup(
        generation: "オンライン", 
        names: [
            "石森 璃花",
            "谷口 愛季",
            "櫻坂46"
        ]
    ),
    MemberGroup(
        generation: "二期", 
        names: [
            "遠藤 光莉", 
            "大沼 晶保", 
            "大園 玲",
            "幸阪 茉里乃",
            "武元 唯衣",
            "田村 保乃",
            "藤吉 夏鈴",
            "増本 綺良",
            "松田 里奈",
            "森田 ひかる",
            "守屋 麗奈",
            "山﨑 天"
        ]
    ),
    MemberGroup(
        generation: "三期", 
        names: [
            "遠藤 理子",
            "小田倉 麗奈",
            "小島 凪紗",
            "中嶋 優月",
            "的野 美青",
            "向井 純葉",
            "村井 優",
            "村山 美羽",
            "山下 瞳月"
        ]
    ),
    MemberGroup(
        generation: "四期", 
        names: [
            "浅井 琥珀",
            "稲熊 陽菜",
            "勝又 陽",
            "佐藤 寧音",
            "中川 千尋",
            "松本 和子",
            "目黒 緋彩",
            "山川 結衣",
            "山田 桃実"
        ]
    ),
]


struct TalkTabView: View {
    let columns = [
        GridItem(.flexible()),
        GridItem(.flexible()),
        GridItem(.flexible()),
    ]
    var body: some View {
        HeaderView(title: "トーク", icons: true, isBlog: false, isSubpage: false)

        ScrollView {
            LazyVGrid(columns: columns, spacing: 16) {
                ForEach(memberData, id: \.generation) { memberGroup in
                    Section(
                        header: GenerationHeader(title: memberGroup.generation),
                        footer: Group {
                            if memberGroup.generation != memberData.last?.generation {
                                Rectangle()
                                    .fill(Color.gray.opacity(0.2))
                                    .frame(height: 2)
                                    .padding(.horizontal, 20)
                            }
                        }
                        ) {
                        ForEach(memberGroup.names, id: \.self) { name in
                            // Placeholder for member image
                            VStack {
                                Circle()
                                    .fill(Color.gray.opacity(0.3))
                                    .frame(width: 96, height: 96)
                                Text(name)
                            }
                            .foregroundColor(sakuraPink)
                            .padding(.bottom, 16)
                        }
                    }
                }
            }
        }
        .scrollIndicators(.hidden)
    }
}

// A helper view to style generation headers
struct GenerationHeader: View {
    let title: String
    
    var body: some View {
        HStack {
            Text(title)
                .font(.headline)
                .padding(.horizontal, 16)
                .padding(.vertical, 16)
                .background(Color.white.opacity(0.9))
                .foregroundColor(sakuraPink) 
        }
        .frame(maxWidth: .infinity)
        .background(Color.white) 
    }
}

#Preview {
    TalkTabView()
}