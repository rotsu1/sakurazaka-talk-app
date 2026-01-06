//
//  sakurazaka_talk_appApp.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 3/1/2026.
//

import SwiftUI
import SwiftData

@main
struct sakurazaka_talk_appApp: App {
    var body: some Scene {
        WindowGroup {
            ContentView()
                .modelContainer(for: Member.self)
        }
    }
}
