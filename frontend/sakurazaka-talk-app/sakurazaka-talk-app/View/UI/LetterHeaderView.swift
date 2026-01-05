//
//  LetterHeaderView.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 5/1/2026.
//

import SwiftUI

struct LetterHeaderView: View {
    let isEdit: Bool
    @Environment(\.dismiss) var dismiss

    var body: some View {
        VStack(spacing: 0) {
            ZStack {
                Text(isEdit ? "編集" : "レター確認")
                    .font(.headline)
                    .fontWeight(.semibold)
                    .foregroundColor(sakuraPink)

                HStack {
                    Button(action: {
                        dismiss()
                    }) {
                        Image(systemName: "chevron.left")
                            .font(.system(size: 22, weight: .light))
                    }

                    Spacer()

                    if isEdit {
                        HStack(spacing: 20) {
                            NavigationLink(destination: LetterConfirmView()) {
                                Text("完了")
                                    .font(.system(size: 22, weight: .light))
                            }
                        }
                    }
                }
                .foregroundColor(sakuraPink)
                .padding(.horizontal, 16)
            }
            .padding(.bottom, 12)
            .padding(.top, 8)

            Rectangle()
                .fill(
                    LinearGradient(
                        colors: [
                            Color.white.opacity(0.5),
                            Color.rgb(red: 213, green: 0, blue: 255)
                        ],
                        startPoint: .leading,
                        endPoint: .trailing
                    )
                )
                .frame(height: 2)
        }
        .background(Color.white)
    }
}