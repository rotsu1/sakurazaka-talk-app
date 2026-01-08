//
//  BlogTabView.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 4/1/2026.
//

import SwiftUI
import SwiftData

struct BlogTabView: View {
    @Query(sort: \Blog.createdAt, order: .reverse) private var blogs: [Blog]
    @Environment(\.modelContext) private var modelContext

    var body: some View {
        VStack {
            HeaderView(title: "ブログ", icons: false, isBlog: true, isSubpage: false)
            ScrollView {
                LazyVStack(spacing: 16) {
                    ForEach(blogs) { blog in
                        BlogItemView(blog: blog)
                            .padding(.horizontal, 16)
                    }
                }
            }
        }
        .task {
            let service = BlogService(modelContext: modelContext)
            do {
                try await service.syncBlogs()
            } catch {
                print("❌ unkoSync failed: \(error)") // This will print the actual reason
            }
        }
    }
}

struct BlogItemView: View {
    let blog: Blog

    var body: some View {
        NavigationLink(destination: BlogView(blog: blog)) {
            HStack {
                Rectangle()
                    .fill(Color.gray.opacity(0.3))
                    .frame(width: 96, height: 96)
                    .cornerRadius(4)
                VStack(alignment: .leading, spacing: 0) {
                    Text(blog.title)
                        .font(.system(size: 18, weight: .medium))
                        .foregroundColor(sakuraPink)

                    Text(blog.content)
                        .font(.system(size: 14, weight: .regular))
                        .foregroundColor(Color(white: 0.6))
                        .lineLimit(2)
                        .multilineTextAlignment(.leading)
                        .padding(.top, 8)
                    
                    Spacer()

                    HStack {
                        Spacer()
                        Text(blog.author)
                            .font(.system(size: 14, weight: .regular))
                            .foregroundColor(sakuraPink)
                        Text(formatterDetailed.string(from: blog.createdAt))
                            .font(.system(size: 14, weight: .regular))
                            .foregroundColor(Color(white: 0.4))
                    }
                    .frame(alignment: .trailing)
                }
                .frame(maxWidth: .infinity)
            }
        }
        .frame(maxWidth: .infinity, alignment: .leading)
        .padding(8)
        .background(Color(white: 0.97), in: RoundedRectangle(cornerRadius: 4))
    }
}

#Preview {
    BlogTabView()
}