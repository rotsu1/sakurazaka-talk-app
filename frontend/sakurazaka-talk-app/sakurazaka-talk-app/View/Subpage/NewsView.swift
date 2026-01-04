//
//  NewsView.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 4/1/2026.
//

import SwiftUI

struct NewsView: View {
    let content: String

    var body: some View {
        HeaderView(title: "ニュース", icons: false, isBlog: false, isSubpage: true)
        
        ScrollView {
            
        }
        .navigationBarHidden(true) 
        .navigationBarBackButtonHidden(true)
    }
}
