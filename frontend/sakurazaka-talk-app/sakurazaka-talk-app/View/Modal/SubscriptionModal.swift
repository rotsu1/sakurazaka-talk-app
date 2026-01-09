//
//  SubscriptionModal.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 6/1/2026.
//

import SwiftUI
import SwiftData

struct SubscriptionModal: View {
    @Environment(\.modelContext) private var modelContext
    @Binding var isPresented: Bool
    var member: Member?
    var service: SubscriptionService

    var body: some View {
        ZStack {
            Color.black.opacity(0.4)
                .edgesIgnoringSafeArea(.all)
                .onTapGesture {
                    isPresented = false
                }
            
            VStack(spacing: 0) {
                Image(member?.avatarUrl ?? "profile_image")
                    .resizable()
                    .aspectRatio(contentMode: .fill)
                    .frame(width: 120, height: 120)
                    .clipShape(Circle())
                    .padding(.top, 30)
                
                Text(member?.name ?? "")
                    .font(.system(size: 20, weight: .bold))
                    .padding(.top, 15)
                
                Text("定期購入をはじめますか？")
                    .font(.system(size: 16))
                    .foregroundColor(.gray)
                    .padding(.top, 8)
                
                VStack(spacing: 0) {
                    Text("月額")
                        .font(.system(size: 14))
                        .frame(maxWidth: .infinity)
                        .padding(.vertical, 8)
                        .background(Color(red: 0.9, green: 0.6, blue: 0.7))
                        .foregroundColor(.white)
                    
                    HStack(alignment: .lastTextBaseline) {
                        Text("¥")
                            .font(.system(size: 14))
                        Text("300")
                            .font(.system(size: 24, weight: .bold))
                        Text("/ 月")
                            .font(.system(size: 14))
                    }
                    .padding(.vertical, 20)
                }
                .overlay(RoundedRectangle(cornerRadius: 8).stroke(Color(red: 0.9, green: 0.6, blue: 0.7), lineWidth: 1))
                .cornerRadius(8)
                .padding(.horizontal, 25)
                .padding(.top, 25)
                
                VStack(alignment: .leading, spacing: 5) {
                    Text("注意事項")
                        .font(.system(size: 14, weight: .bold))
                        .padding(.bottom, 2)
                    
                    Text("• ご利用のApple IDアカウントに課金されます。")
                    Text("• 定期購入の期間終了日の24時間以上前に自動更新を解除しない限り、定期購入の期間が自動更新されます。")
                }
                .font(.system(size: 11))
                .foregroundColor(.gray)
                .padding(.horizontal, 25)
                .padding(.top, 20)
                
                Spacer().frame(height: 30)
                
                Divider()
                
                Button(action: {
                    Task {
                        do {
                            try await service.subscribe(to: member!, context: modelContext)
                            isPresented = false // Close modal on success
                        } catch {
                            print("Purchase failed: \(error)")
                        }
                    }
                }) {
                    Text("定期購入をはじめる")
                        .fontWeight(.medium)
                        .foregroundColor(Color(red: 0.9, green: 0.5, blue: 0.6))
                        .frame(maxWidth: .infinity, minHeight: 50)
                }
                
                Divider()
                
                Button(action: { isPresented = false }) {
                    Text("キャンセル")
                        .foregroundColor(.black)
                        .frame(maxWidth: .infinity, minHeight: 50)
                }
            }
            .background(Color.white)
            .cornerRadius(15)
            .padding(.horizontal, 30)
            .shadow(radius: 10)
        }
    }
}
