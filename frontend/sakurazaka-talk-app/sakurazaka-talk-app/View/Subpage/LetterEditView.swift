//
//  LetterEditView.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 5/1/2026.
//

import SwiftUI

struct LetterEditView: View {
    @State private var message: String = ""

    var body: some View {
        VStack(spacing: 0) {
            LetterHeaderView(isEdit: true)

            Spacer()

            ZStack {
                Image(systemName: "photo")
                    .resizable()
                    .frame(width: 350, height: 500)
                    .foregroundColor(.gray.opacity(0.5))
                    .scaledToFit()
                    .padding(.horizontal, 20)

                ZStack {
                    RoundedRectangle(cornerRadius: 12)
                        .fill(Color.black.opacity(0.25))
                    
                    if message.isEmpty {
                        VStack(spacing: 12) {
                            Image(systemName: "plus")
                                .font(.system(size: 40, weight: .thin))
                            Text("メッセージを入力")
                                .font(.body)
                        }
                        .foregroundColor(.white)
                    } else {
                        Text(message)
                            .foregroundColor(.white)
                            .font(.body)
                            .multilineTextAlignment(.center)
                            .padding()
                    }
                }
                .frame(width: 260, height: 180)
                .onTapGesture {
                    withAnimation {
                        if message.isEmpty {
                            message = "お誕生日おめでとう！\n素敵な一年になりますように。"
                        } else {
                            message = ""
                        }
                    }
                }
            }
            
            Spacer()
        }
        .background(Color.white)
        .navigationBarHidden(true)
        .navigationBarBackButtonHidden(true)
    }
}

#Preview {
    LetterEditView()
}
