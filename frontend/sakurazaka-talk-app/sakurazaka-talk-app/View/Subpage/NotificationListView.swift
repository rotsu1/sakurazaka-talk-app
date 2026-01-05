//
//  NotificationListView.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 3/1/2026.
//

import SwiftUI

struct NotificationItem: Identifiable, Hashable {
    let id = UUID()
    let title: String
    let content: String
    let timestamp: Date
}

let calendar = Calendar.current

let notificationItems = [
    // 1. Rena Moriya Birthday Notification
    NotificationItem(
        title: "守屋麗奈のバースデーレターテンプレートを公開！", 
        content: """
        明日1月2日(金)は守屋麗奈の誕生日🎂
        誕生日を記念して、本人作成のバースデーレターテンプレートを、1月23日(金)までの期間限定で公開中です！
        ぜひチェックしてください！
        """, 
        timestamp: calendar.date(from: DateComponents(year: 2026, month: 1, day: 1, hour: 18, minute: 5))!
    ),
    
    // 2. New Year Template Notification
    NotificationItem(
        title: "2026年お正月レター＆カードテンプレートを公開いたしました！", 
        content: """
        現在実施中の「Thanks 2025 / Hello 2026 レターキャンペーン」の対象テンプレートに追加され、2026年1月31日(土)までの期間限定公開です。

        「Thanks 2025 / Hello 2026 レターキャンペーン」では、対象のテンプレートを使用して期間中にレターやカードを2通以上送ると、送ったメンバーの直筆サイン＆宛名入りリアルレターが当たるチャンス！

        この一年を振り返るメッセージや新たな一年の挨拶をレターで送ってみてください！

        【キャンペーン期間】
        2025年12月25日(木)12:00～2026年1月7日(水)23:59

        【景品】
        直筆サイン＆宛名入りリアルレター
        メンバー32名×各2名様（計64名様）

        ※エントリーの際、ご希望の宛名の入力と景品希望メンバーを選択いただきます。
        """, 
        timestamp: calendar.date(from: DateComponents(year: 2025, month: 12, day: 31))!
    ),
    
    // 3. Campaign Announcement
    NotificationItem(
        title: "「Thanks 2025 / Hello 2026 レターキャンペーン」を実施いたします！", 
        content: """
        櫻坂46メッセージでは「Thanks 2025 / Hello 2026 レターキャンペーン」を実施いたします！
        対象のテンプレートを使用して期間中にレターやカードを2通以上送ると、送ったメンバーの直筆サイン＆宛名入りリアルレターが当たるチャンス！

        この一年を振り返るメッセージや新たな一年の挨拶をレターで送ってみてください！

        【キャンペーン期間】
        2025年12月25日(木)12:00～2026年1月7日(水)23:59

        【景品】
        直筆サイン＆宛名入りリアルレター
        メンバー32名×各2名様（計64名様）

        ※エントリーの際、ご希望の宛名の入力と景品希望メンバーを選択いただきます。

        【参加条件】
        ①以下エントリーページから2026年1月7日(水)23:59までにエントリーを完了
        https://sakurazaka46.com/s/s46/form/question?cd=msgappcp251225&uid=13dd0157-85e1-4e12-9313-fe931fd251c3

        ②2025年12月25日(木)12:00～2026年1月7日(水)23:59の間で、希望の対象メンバーに2通以上対象のテンプレートでレターやカードを送付していること
        """, 
        timestamp: calendar.date(from: DateComponents(year: 2025, month: 12, day: 25))!
    )
]

struct NotificationListView: View {
    let calendar = Calendar.current

    var body: some View {
        VStack {
            HeaderView(title: "お知らせ", icons: false, isBlog: false, isSubpage: true)
            
            ScrollView {
                LazyVStack(alignment: .leading, spacing: 12) {
                    ForEach(notificationItems) { item in
                        NavigationLink(
                            destination: NotificationView(notificationItem: item)
                        ) {
                            VStack(alignment: .leading, spacing: 8) {
                                let dateString = formatterSimple.string(
                                    from: item.timestamp
                                )
                                Text(dateString)
                                    .foregroundColor(Color(white: 0.6))
                                    .font(.system(size: 13, weight: .regular))
                                    .lineLimit(1)
                                Text(item.title)
                                    .foregroundColor(Color(white: 0.5))
                                    .font(.system(size: 17, weight: .regular))
                                    .lineLimit(1)
                                Text(item.content)
                                    .foregroundColor(Color(white: 0.3))
                                    .font(.system(size: 14, weight: .medium))
                                    .lineLimit(1)
                            }
                            .padding()
                            .frame(maxWidth: .infinity, alignment: .leading)
                            .background(Color.rgb(red: 247, green: 247, blue: 247))
                        }
                    }
                }
                .padding()
            }
            .navigationBarHidden(true) 
            .navigationBarBackButtonHidden(true)
        }
    }
}

#Preview {
    NotificationListView()
}