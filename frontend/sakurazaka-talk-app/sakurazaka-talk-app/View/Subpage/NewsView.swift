//
//  NewsView.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 4/1/2026.
//

import SwiftUI
import WebKit

struct NewsView: View {
    let content: String

    var body: some View {
        VStack {
            HeaderView(title: "ニュース", icons: false, isBlog: false, isSubpage: true)

            NewsWebView(content: content)
                .navigationBarHidden(true) 
                .navigationBarBackButtonHidden(true)
        }
    }
}

struct NewsWebView: UIViewRepresentable {
    let content: String

    func makeUIView(context: Context) -> WKWebView {
        return WKWebView()
    }

    func updateUIView(_ uiView: WKWebView, context: Context) {
        uiView.loadHTMLString(content, baseURL: nil)
    }
}