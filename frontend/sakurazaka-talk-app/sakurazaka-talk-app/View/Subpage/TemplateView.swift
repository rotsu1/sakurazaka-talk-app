//
//  TemplateView.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 5/1/2026.
//

import SwiftUI

struct TemplateView: View {
    @State private var selectedTab: TemplateTab = .letter
    
    let sakuraPink = Color(red: 242/255, green: 139/255, blue: 169/255)
    
    let columns = [
        GridItem(.flexible(), spacing: 15),
        GridItem(.flexible(), spacing: 15)
    ]

    var body: some View {
        VStack(spacing: 0) {
            HeaderView(title: "テンプレートを選択", icons: false, isBlog: false, isSubpage: true)

            TemplateTabBar(selectedTab: $selectedTab)
                .padding(.top, 5)

            if selectedTab == .draft {
                EmptyStateView(sakuraPink: sakuraPink)
            } else {
                ScrollView {
                    LazyVGrid(columns: columns, spacing: 15) {
                        ForEach(0..<getItemCount(for: selectedTab), id: \.self) { index in
                            TemplateCard(index: index)
                        }
                    }
                    .padding(15)
                }
            }
        }
        .navigationBarHidden(true)
        .background(Color.white)
    }
    
    @Namespace private var namespace
    
    func getItemCount(for tab: TemplateTab) -> Int {
        switch tab {
        case .letter: return 6
        case .card: return 4
        case .draft: return 0
        }
    }
}

struct EmptyStateView: View {
    let sakuraPink: Color
    
    var body: some View {
        VStack(spacing: 24) {
            Spacer()
            
            Image(systemName: "envelope")
                .resizable()
                .scaledToFit()
                .frame(width: 60, height: 60)
                .foregroundColor(Color(UIColor.systemGray4))
            
            Text("下書きのレターがありません")
                .font(.system(size: 15))
                .foregroundColor(.black)
                .frame(alignment: .center)
            
            Text("編集中のレターが表示されます。")
                .font(.system(size: 14))
                .foregroundColor(Color(white: 0.5))
                .frame(alignment: .center)
            
            Spacer()
            Spacer()
        }
        .frame(maxWidth: .infinity, maxHeight: .infinity)
        .background(Color.white)
    }
}

enum TemplateTab: String, CaseIterable {
    case letter = "レター"
    case card = "カード"
    case draft = "下書き"
}

struct TemplateCard: View {
    let index: Int
    
    var body: some View {
        GeometryReader { geo in
            NavigationLink(destination: LetterEditView()) {
                ZStack {
                    Color(hue: Double(index) * 0.1, saturation: 0.1, brightness: 0.95)
                    
                    VStack {
                        Image(systemName: "photo")
                            .font(.largeTitle)
                            .foregroundColor(.gray.opacity(0.5))
                        Text("Template \(index + 1)")
                            .font(.caption)
                            .foregroundColor(.gray)
                    }
                }
                .cornerRadius(4)
                .shadow(color: Color.black.opacity(0.1), radius: 2, x: 0, y: 1)
            }
        }
        .aspectRatio(0.65, contentMode: .fit)
        .navigationBarHidden(true)
        .navigationBarBackButtonHidden(true)
    }
}

struct TemplateTabBar: View {
    @Binding var selectedTab: TemplateTab
    
    @Namespace private var namespace
    
    let sakuraPink = Color(red: 242/255, green: 139/255, blue: 169/255)
    
    var body: some View {
        HStack(spacing: 0) {
            ForEach(TemplateTab.allCases, id: \.self) { tab in
                Button(action: {
                    withAnimation(.easeInOut(duration: 0.2)) {
                        selectedTab = tab
                    }
                }) {
                    VStack(spacing: 12) {
                        Text(tab.rawValue)
                            .font(.subheadline)
                            .foregroundColor(selectedTab == tab ? sakuraPink : .gray)
                        
                        ZStack(alignment: .bottom) {
                            Rectangle()
                                .fill(Color.white)
                                .frame(height: 1)
                            
                            if selectedTab == tab {
                                Rectangle()
                                    .fill(sakuraPink)
                                    .frame(height: 3)
                                    .matchedGeometryEffect(id: "TabUnderline", in: namespace)
                            }
                        }
                        .frame(height: 3)
                    }
                }
                .buttonStyle(PlainButtonStyle())
            }
        }
        .frame(height: 54)
    }
}

#Preview {
    TemplateView()
}
