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