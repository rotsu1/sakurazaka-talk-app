//
//  TalkView.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 5/1/2026.
//

import SwiftUI
import SwiftData
import AVKit

struct TalkView: View {
    let member: Member
    @Query(sort: \Message.createdAt) var messages: [Message] // Assumes you use the @Query init pattern we discussed
    @Environment(\.modelContext) private var modelContext
    
    let themePink = Color(red: 235/255, green: 130/255, blue: 155/255)

    var body: some View {
        VStack {
            TalkHeaderView(title: member.name)
            ScrollView {
                VStack(alignment: .leading, spacing: 20) {
                    ForEach(messages) { message in
                        ChatRow(avatarName: "person.crop.circle.fill") {
                            messageContent(for: message)
                        }
                    }
                }
                .padding(.top, 20)
                .padding(.bottom, 50)
            }
        }
        .navigationBarHidden(true)
        .task {
            let service = MessageService(modelContext: modelContext)
            do {
                try await service.syncMessages()
            } catch {
                print("❌ Sync failed: \(error)") // This will print the actual reason
            }
        }
    }

    @ViewBuilder
    func messageContent(for message: Message) -> some View {
        switch message.type {
        case "text":
            Text(message.content)
                .font(.system(size: 15))
                .foregroundColor(Color(UIColor.darkGray))
        
        case "image":
            if let data = message.data, let image = UIImage(data: data) {
                Image(uiImage: image)
                    .resizable()
                    .aspectRatio(contentMode: .fill)
                    .frame(width: 200, height: 250)
                    .clipped()
            } else {
                placeholderMediaView(icon: "photo")
            }

        case "video":
            if let data = message.data {
                VideoPlayerBubble(data: data)
            } else {
                placeholderMediaView(icon: "video.fill")
            }

        case "voice":
            if let data = message.data {
                AudioPlayerBubble(data: data, themeColor: themePink)
            } else {
                placeholderMediaView(icon: "waveform")
            }
            
        default:
            Text(message.content)
        }
    }

    func placeholderMediaView(icon: String) -> some View {
        ZStack {
            Color.gray.opacity(0.3)
            Image(systemName: icon)
                .foregroundColor(.white)
        }
        .frame(width: 200, height: 250)
        .cornerRadius(12)
    }
}

struct VideoPlayerBubble: View {
    let data: Data
    @State private var playerURL: URL?

    var body: some View {
        Group {
            if let url = playerURL {
                VideoPlayer(player: AVPlayer(url: url))
                    .frame(width: 200, height: 250)
                    .cornerRadius(12)
            } else {
                ProgressView()
                    .frame(width: 200, height: 250)
            }
        }
        .onAppear {
            let tempDir = FileManager.default.temporaryDirectory
            let fileURL = tempDir.appendingPathComponent(UUID().uuidString + ".mp4")
            try? data.write(to: fileURL)
            self.playerURL = fileURL
        }
    }
}

struct AudioPlayerBubble: View {
    let data: Data
    var themeColor: Color
    @State private var audioPlayer: AVAudioPlayer?
    @State private var isPlaying = false

    var body: some View {
        VStack(spacing: 15) {
            // ... (Your existing GeometryReader progress bar code here)
            HStack {
                Image(systemName: "speaker.wave.2.fill").foregroundColor(.gray)
                Spacer()
                Button {
                    togglePlay()
                } label: {
                    Image(systemName: isPlaying ? "pause.fill" : "play.fill")
                        .foregroundColor(themeColor)
                        .font(.system(size: 24))
                }
                Spacer()
                Text(formatDuration(audioPlayer?.duration ?? 0))
                    .font(.caption).foregroundColor(.gray)
            }
        }
        .frame(width: 200)
        .onAppear { audioPlayer = try? AVAudioPlayer(data: data) }
    }

    func togglePlay() {
        guard let player = audioPlayer else { return }
        if player.isPlaying { player.pause() } else { player.play() }
        isPlaying = player.isPlaying
    }
    
    func formatDuration(_ d: TimeInterval) -> String {
        let min = Int(d) / 60
        let sec = Int(d) % 60
        return String(format: "%02d:%02d", min, sec)
    }
}

// import SwiftUI

// struct TalkView: View {
//     let member: Member

//     let themePink = Color(red: 235/255, green: 130/255, blue: 155/255)
//     let bgGray = Color(red: 245/255, green: 245/255, blue: 245/255)

//     var body: some View {
//         VStack {
//             TalkHeaderView(title: "石森 璃花")
//             ScrollView {
//                 VStack(alignment: .leading, spacing: 20) {
//                     let sortedMessages = messages.sorted(by: { $0.createdAt < $1.createdAt })
//                     ForEach(sortedMessages) { message in
//                         if message.type == "text" {
//                             ChatRow(avatarName: "person.crop.circle.fill") {
//                                 Text(message.content)
//                             }
//                         }
//                         else if message.type == "image" {
//                             ChatRow(avatarName: "person.crop.circle.fill") {
//                                 if let data = message.data, let image = UIImage(data: data) {
//                                     Image(uiImage: image)
//                                         .resizable()
//                                         .aspectRatio(contentMode: .fill)
//                                         .frame(width: 200, height: 250)
//                                         .background(Color.gray.opacity(0.3))
//                                         .clipped()
//                                 }
//                             }
//                         }
//                         else if message.type == "voice" {
//                             ChatRow(avatarName: "person.crop.circle.fill") {
//                                 AudioPlayerBubble(themeColor: themePink)
//                             }
//                         }
//                     }
                    
//                     ChatRow(avatarName: "person.crop.circle.fill") {
//                         Image(systemName: "photo")
//                             .resizable()
//                             .aspectRatio(contentMode: .fill)
//                             .frame(width: 200, height: 250)
//                             .background(Color.gray.opacity(0.3))
//                             .clipped()
//                     }

//                     ChatRow(avatarName: "person.crop.circle.fill") {
//                         Text("大好きです！")
//                             .font(.system(size: 15))
//                             .foregroundColor(Color(UIColor.darkGray))
//                     }

//                     ChatRow(avatarName: "person.crop.circle.fill") {
//                         VStack(alignment: .leading, spacing: 8) {
//                             Text("港区パセリが好きです！MVはかっこいいし、ライブの時に胸を撃ち抜かれました。")
//                                 .font(.system(size: 15))
//                                 .foregroundColor(Color(UIColor.darkGray))
                            
//                             Text("君のことを思いながらも好きです！")
//                                 .font(.system(size: 15))
//                                 .foregroundColor(Color(UIColor.darkGray))
//                                 .lineSpacing(4)
//                         }
//                     }

//                     ChatRow(avatarName: "person.crop.circle.fill") {
//                         AudioPlayerBubble(themeColor: themePink)
//                     }

//                 }
//                 .padding(.top, 20)
//                 .padding(.bottom, 50)
//             }
//         }
//         .navigationBarHidden(true)
//         .navigationBarBackButtonHidden(true)
//     }
// }

struct ChatRow<Content: View>: View {
    var avatarName: String
    let content: Content

    init(avatarName: String, @ViewBuilder content: () -> Content) {
        self.avatarName = avatarName
        self.content = content()
    }

    var body: some View {
        HStack(alignment: .top, spacing: 10) {
            Image(uiImage: UIImage(named: "avatar_placeholder") ?? UIImage(systemName: "person.circle")!) 
                .resizable()
                .scaledToFill()
                .frame(width: 40, height: 40)
                .clipShape(Circle())
                .overlay(Circle().stroke(Color.gray.opacity(0.1), lineWidth: 1))
                .foregroundColor(.gray)

            content
                .padding(12)
                .background(Color(white: 0.95))
                .cornerRadius(12)
                .shadow(color: Color.black.opacity(0.05), radius: 2, x: 0, y: 1)
        }
        .padding(.horizontal, 8)
    }
}

// struct AudioPlayerBubble: View {
//     var themeColor: Color
    
//     var body: some View {
//         VStack(spacing: 15) {
//             GeometryReader { geometry in
//                 ZStack(alignment: .leading) {
//                     Capsule()
//                         .fill(Color.gray.opacity(0.2))
//                         .frame(height: 4)
                    
//                     Capsule()
//                         .fill(themeColor.opacity(0.6))
//                         .frame(width: geometry.size.width * 0.1, height: 4)
                    
//                     Circle()
//                         .fill(themeColor)
//                         .frame(width: 14, height: 14)
//                         .offset(x: geometry.size.width * 0.1 - 7)
//                 }
//             }
//             .frame(height: 14)

//             HStack {
//                 Image(systemName: "speaker.wave.2.fill")
//                     .foregroundColor(.gray)
//                     .font(.system(size: 16))
                
//                 Spacer()
                
//                 Image(systemName: "play.fill")
//                     .foregroundColor(themeColor)
//                     .font(.system(size: 24))
                
//                 Spacer()
                
//                 Text("00:47")
//                     .font(.caption)
//                     .foregroundColor(.gray)
//             }
//         }
//         .frame(width: 200)
//         .padding(.vertical, 4)
//     }
// }

// #Preview {
//     TalkView()
// }
