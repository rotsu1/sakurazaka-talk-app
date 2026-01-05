//
//  TalkView.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 5/1/2026.
//

import SwiftUI

struct TalkView: View {
    let themePink = Color(red: 235/255, green: 130/255, blue: 155/255)
    let bgGray = Color(red: 245/255, green: 245/255, blue: 245/255)

    var body: some View {
        VStack {
            TalkHeaderView(title: "石森 璃花")
            ScrollView {
                VStack(alignment: .leading, spacing: 20) {
                    
                    ChatRow(avatarName: "person.crop.circle.fill") {
                        Image(systemName: "photo")
                            .resizable()
                            .aspectRatio(contentMode: .fill)
                            .frame(width: 200, height: 250)
                            .background(Color.gray.opacity(0.3))
                            .clipped()
                    }

                    ChatRow(avatarName: "person.crop.circle.fill") {
                        Text("大好きです！")
                            .font(.system(size: 15))
                            .foregroundColor(Color(UIColor.darkGray))
                    }

                    ChatRow(avatarName: "person.crop.circle.fill") {
                        VStack(alignment: .leading, spacing: 8) {
                            Text("港区パセリが好きです！MVはかっこいいし、ライブの時に胸を撃ち抜かれました。")
                                .font(.system(size: 15))
                                .foregroundColor(Color(UIColor.darkGray))
                            
                            Text("君のことを思いながらも好きです！")
                                .font(.system(size: 15))
                                .foregroundColor(Color(UIColor.darkGray))
                                .lineSpacing(4)
                        }
                    }

                    ChatRow(avatarName: "person.crop.circle.fill") {
                        AudioPlayerBubble(themeColor: themePink)
                    }

                }
                .padding(.top, 20)
                .padding(.bottom, 50)
            }
        }
        .navigationBarHidden(true)
        .navigationBarBackButtonHidden(true)
    }
}

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

struct AudioPlayerBubble: View {
    var themeColor: Color
    
    var body: some View {
        VStack(spacing: 15) {
            GeometryReader { geometry in
                ZStack(alignment: .leading) {
                    Capsule()
                        .fill(Color.gray.opacity(0.2))
                        .frame(height: 4)
                    
                    Capsule()
                        .fill(themeColor.opacity(0.6))
                        .frame(width: geometry.size.width * 0.1, height: 4)
                    
                    Circle()
                        .fill(themeColor)
                        .frame(width: 14, height: 14)
                        .offset(x: geometry.size.width * 0.1 - 7)
                }
            }
            .frame(height: 14)

            HStack {
                Image(systemName: "speaker.wave.2.fill")
                    .foregroundColor(.gray)
                    .font(.system(size: 16))
                
                Spacer()
                
                Image(systemName: "play.fill")
                    .foregroundColor(themeColor)
                    .font(.system(size: 24))
                
                Spacer()
                
                Text("00:47")
                    .font(.caption)
                    .foregroundColor(.gray)
            }
        }
        .frame(width: 200)
        .padding(.vertical, 4)
    }
}

#Preview {
    TalkView()
}
