//
//  BlogView.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 4/1/2026.
//

import SwiftUI

struct BlogView: View {
    let blog: Blog

    var body: some View {
        VStack {
            HeaderView(title: blog.member, icons: false, isBlog: true, isSubpage: true)
            
            ScrollView {
                LazyVStack(spacing: 16) {
                    BlogHeaderView(date: blog.createdAt, title: blog.title)
                        .padding()
                        .padding(.horizontal, 16)
                    Text(blog.content)
                        .font(.system(size: 16, weight: .regular))
                        .foregroundColor(Color(white: 0.4))
                        .frame(maxWidth: .infinity, alignment: .leading)
                        .padding()
                        .padding(.horizontal, 16)
                }
            }
            .navigationBarHidden(true) 
            .navigationBarBackButtonHidden(true)
        }
    }
}

struct BlogHeaderView: View {
    let date: Date
    let title: String

    var body: some View {
        let year = calendar.component(.year, from: date)
        let month = calendar.component(.month, from: date)
        let day = calendar.component(.day, from: date)

        HStack {
            ZStack {
                VStack {
                    HStack {
                        VStack(alignment: .leading) {
                            Text(String(year))
                                .font(.system(size: 12, weight: .regular))
                                .foregroundColor(Color(white: 0.5))
                            Text(String(month))
                                .font(.system(size: 10, weight: .regular))
                                .foregroundColor(Color(white: 0.5))
                        }
                        Spacer()
                    }
                    .padding(4)
                    .offset(y: 4)
                    Spacer()
                    HStack {
                        Spacer()
                        Text(String(day))
                            .font(.system(size: 24, weight: .regular))
                            .foregroundColor(Color(white: 0.3))
                            .padding(.trailing, 8)
                            .padding(.bottom, 8)
                    }
                }
            }
            .frame(width: 64, height: 64)
            .overlay(
                Rectangle().frame(width: 1.5).foregroundColor(Color(white: 0.6)),
                alignment: .trailing
            )
            .overlay(
                Rectangle().frame(height: 1.5).foregroundColor(Color(white: 0.6)),
                alignment: .bottom
            )
            .overlay(
                DiagonalLine()
                    .stroke(Color(white: 0.7), lineWidth: 1)
            )

            Text(title)
                .font(.system(size: 16, weight: .regular))
                .foregroundColor(Color(white: 0.4))
                .padding(.leading, 12)
                .frame(maxWidth: .infinity, alignment: .leading)
                .frame(height: 64)
                .overlay(
                    Rectangle().frame(height: 1.5).foregroundColor(Color(white: 0.6)),
                    alignment: .bottom
                )

            Spacer()
        }
        .frame(maxWidth: .infinity)

    }
}

struct DiagonalLine: Shape {
    func path(in rect: CGRect) -> Path {
        var path = Path()
        // Start Top-Right
        path.move(to: CGPoint(x: rect.maxX - 8, y: 8))
        // Draw to Bottom-Left
        path.addLine(to: CGPoint(x: 8, y: rect.maxY - 8))
        return path
    }
}

#Preview {
    BlogView(blog: Blog(title: "Title", content: "Content", member: "Sugai", createdAt: Date()))
}
