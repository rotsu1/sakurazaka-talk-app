//
//  BlogTabView.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 4/1/2026.
//

import SwiftUI

struct Blog: Identifiable {
    let id = UUID()
    let title: String
    let content: String
    let member: String
    let createdAt: Date
}

let blogData = [
    Blog(
        title: "#370 Instagram始めます！", 
        content: """
        こんばんは！




        櫻坂46　二期生　宮城県出身
        　　まつりちゃんこと松田里奈です。





        ブログを開いていただきありがとうございます。






        お知らせがあります！






        わたし！！
        """, 
        member: "松田 里奈",
        createdAt: calendar.date(from: DateComponents(year: 2026, month: 1, day: 1, hour: 21, minute: 0))!
    ),
    Blog(
        title: "#369 2026年スタート！", 
        content: """
        明けましておめでとうございます！

        櫻坂46　二期生　宮崎県出身
        　　まつりちゃんこと松田里奈です。



        ブログ開いていただきありがとうございます。





        年末年始も放送しているTHE TIME,! 私も元旦から出演させていただきました！



        私は東京スカイツリーから初日の出中継！
        """, 
        member: "松田 里奈",
        createdAt: calendar.date(from: DateComponents(year: 2026, month: 1, day: 1, hour: 18, minute:42))!
    ),
    Blog(
        title: "2026", 
        content: """
        ブログを開いてくださり、ありがとうございます☺︎


        櫻坂46 四期生 埼玉県出身21歳の浅井恋乃未（あさいこのみ）です。

        ※21歳になったことをすっかり忘れていて… 前回と前々回のブログで20歳と書いてしまっていました( ; ; )






        あけましておめでとうございます！！
        """, 
        member: "浅井 恋乃未",
        createdAt: calendar.date(from: DateComponents(year: 2026, month: 1, day: 1, hour: 16, minute:37))!
    ),
]

struct BlogTabView: View {
    var body: some View {
        HeaderView(title: "ブログ", icons: false, isBlog: true, isSubpage: false)
        ScrollView {
            LazyVStack(spacing: 16) {
                ForEach(blogData) { blog in
                    BlogItemView(blog: blog)
                        .padding(.horizontal, 16)
                }
            }
        }
    }
}

struct BlogItemView: View {
    let blog: Blog

    var body: some View {
        NavigationLink(destination: BlogView(blog: blog)) {
            HStack {
                Rectangle()
                    .fill(Color.gray.opacity(0.3))
                    .frame(width: 96, height: 96)
                    .cornerRadius(4)
                VStack(alignment: .leading, spacing: 0) {
                    Text(blog.title)
                        .font(.system(size: 18, weight: .medium))
                        .foregroundColor(sakuraPink)

                    Text(blog.content)
                        .font(.system(size: 14, weight: .regular))
                        .foregroundColor(Color(white: 0.6))
                        .lineLimit(2)
                        .padding(.top, 8)
                    
                    Spacer()

                    HStack {
                        Spacer()
                        Text(blog.member)
                            .font(.system(size: 14, weight: .regular))
                            .foregroundColor(sakuraPink)
                        Text(formatterDetailed.string(from: blog.createdAt))
                            .font(.system(size: 14, weight: .regular))
                            .foregroundColor(Color(white: 0.4))
                    }
                    .frame(alignment: .trailing)
                }
                .frame(maxWidth: .infinity)
            }
        }
        .frame(maxWidth: .infinity, alignment: .leading)
        .padding(8)
        .background(Color(white: 0.97), in: RoundedRectangle(cornerRadius: 4))
    }
}

#Preview {
    BlogTabView()
}