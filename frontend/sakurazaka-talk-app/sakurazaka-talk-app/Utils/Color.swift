//
//  Color.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 3/1/2026.
//

import SwiftUI

extension Color {
    static func rgb(red: CGFloat, green: CGFloat, blue: CGFloat) -> Color{
        return self.init(red: red/255, green: green/255, blue: blue/255)
    }
    static func rgbo(red: CGFloat, green: CGFloat, blue: CGFloat, opacity: Double) -> Color{
        return self.init(red: red/255, green: green/255, blue: blue/255, opacity: opacity)
    }
}

let sakuraPink = Color.rgb(red: 241, green: 157, blue: 181)

let tagColor: [String: Color] = [
    "メディア": Color.rgb(red: 198, green: 150, blue: 222),
    "ファンクラブ": Color.rgb(red: 124, green: 206, blue: 124),
    "リリース": Color.rgb(red: 240, green: 138, blue: 138),
    "イベント情報": Color.rgb(red: 245, green: 225, blue: 125),
    "グッズ": Color.rgb(red: 238, green: 192, blue: 140),
    "その他": Color.rgb(red: 180, green: 185, blue: 185),
]

let tagTextColor: [String: Color] = [
    "メディア": Color.white,
    "ファンクラブ": Color.white,
    "リリース": Color.white,
    "イベント情報": Color(white: 0.3),
    "グッズ": Color(white: 0.3),
    "その他": Color.white,
]