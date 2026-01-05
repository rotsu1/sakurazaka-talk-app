//
//  LetterConfirmView.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 5/1/2026.
//

import SwiftUI

struct LetterConfirmView: View {    
    // Theme Color
    let sakuraPink = Color(red: 242/255, green: 139/255, blue: 169/255)
    
    var body: some View {
        VStack(spacing: 0) {
            
            LetterHeaderView(isEdit: false)
            
                VStack(spacing: 20) {
                    
                    HStack(spacing: 12) {
                        Image(systemName: "person.crop.circle.fill")
                            .resizable()
                            .foregroundColor(.gray)
                            .frame(width: 40, height: 40)
                            .clipShape(Circle())
                        
                        Text("石森 璃花 さんへ")
                            .font(.system(size: 16))
                            .foregroundColor(.black.opacity(0.8))
                        
                        Spacer()
                    }
                    .padding(.horizontal)
                    .padding(.top, 10)
                    
                    ZStack {
                        Color(UIColor.systemGray6)
                            .aspectRatio(0.65, contentMode: .fit)
                            .overlay(
                                Image(uiImage: UIImage(named: "new_year_template") ?? UIImage())
                                    .resizable()
                                    .scaledToFill()
                            )
                            .clipped()
                        
                        RoundedRectangle(cornerRadius: 8)
                            .fill(Color.gray.opacity(0.4))
                            .frame(width: 220, height: 300)
                            .overlay(
                                VStack(spacing: 10) {
                                    Image(systemName: "plus")
                                        .font(.system(size: 30, weight: .light))
                                    Text("メッセージを入力")
                                        .font(.system(size: 14))
                                }
                                .foregroundColor(.white)
                            )
                    }
                    .padding(.horizontal, 40)
                    
                }
                .padding(.bottom, 20)

            Spacer()
            
            HStack(spacing: 15) {
                Button(action: {
                }) {
                    Text("下書き保存")
                        .font(.system(size: 15))
                        .fontWeight(.medium)
                        .foregroundColor(sakuraPink)
                        .frame(maxWidth: .infinity)
                        .padding(.vertical, 14)
                        .overlay(
                            RoundedRectangle(cornerRadius: 8)
                                .stroke(sakuraPink, lineWidth: 1)
                        )
                }
                
                Button(action: {
                }) {
                    Text("送信")
                        .font(.system(size: 15))
                        .fontWeight(.medium)
                        .foregroundColor(.white)
                        .frame(maxWidth: .infinity)
                        .padding(.vertical, 14)
                        .background(sakuraPink)
                        .cornerRadius(8)
                }
            }
            .padding(.horizontal, 20)
            .padding(.bottom, 10)
            .padding(.top, 10)
        }
        .background(Color.white)
        .navigationBarHidden(true)
        .navigationBarBackButtonHidden(true)
    }
}

#Preview {
    LetterConfirmView()
}
