//
//  ContentView.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 3/1/2026.
//

import SwiftUI

struct ContentView: View {
    var body: some View {
        NavigationStack {
            TabView {
                Tab("Talk", systemImage: "bubble.right") {
                    TalkTabView()
                }
                Tab("Blog", systemImage: "text.page") {
                    BlogTabView()
                }
                Tab("News", systemImage: "megaphone") {
                    NewsTabView()
                }
                Tab("Official", systemImage: "triangleshape") {

                }
                Tab("Fanclub", systemImage: "oval.portrait") {

                }
            }
        }
    }
}

#Preview {
    ContentView()
}
