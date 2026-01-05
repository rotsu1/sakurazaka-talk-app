//
//  FanclubTabView.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 4/1/2026.
//

import SwiftUI

struct FanclubTabView: View {
    @State private var selectedTab: FanclubTabs = .radio

    var body: some View {
        VStack {
            HeaderView(
                title: "櫻坂46ファンクラブ", 
                icons: true, 
                isBlog: false, 
                isSubpage: false
            )
            FanclubTabButtons(selectedTab: $selectedTab)
            ScrollView {
                
            }
            .navigationBarHidden(true) 
            .navigationBarBackButtonHidden(true)
            .padding()
        }
    }
}

enum FanclubTabs: String, CaseIterable, Identifiable {
    case radio = "RADIO"
    case greeting = "GREETING"
    case managersDiary = "MANAGER'S DIARY"
    case history = "HISTORY"
    case costume = "COSTUME"
    case fcMovie = "FC MOVIE"
    case wallpaper = "WALLPAPER"
    case ticket = "TICKET"
    case myPage = "MY PAGE"
    case archive = "ARCHIVE"
    
    var id: String { rawValue }
}

struct FanclubTabButtons: View {
    @Binding var selectedTab: FanclubTabs
    
    @Namespace private var animationNamespace

    var body: some View {
        ScrollViewReader { proxy in
            ScrollView(.horizontal, showsIndicators: false) {
                HStack(spacing: 24) {
                    ForEach(FanclubTabs.allCases) { tab in
                        TabButton(tab: tab)
                            .id(tab)
                    }
                }
                .padding(.horizontal, 16)
            }
            .onChange(of: selectedTab) { newValue, _ in
                withAnimation {
                    proxy.scrollTo(newValue, anchor: .center)
                }
            }
        }
        .frame(height: 44)
    }

    private func TabButton(tab: FanclubTabs) -> some View {
        let isSelected = selectedTab == tab
        
        return Button {
            withAnimation(.spring(response: 0.3, dampingFraction: 0.7)) {
                selectedTab = tab
            }
        } label: {
            VStack(spacing: 6) {
                ZStack {
                    Text(tab.rawValue)
                        .font(.system(size: 16, weight: .regular))
                        .foregroundColor(sakuraPink)
                        .opacity(isSelected ? 0 : 1)
                    
                    Text(tab.rawValue)
                        .font(.system(size: 16, weight: .semibold))
                        .foregroundColor(sakuraPink)
                        .opacity(isSelected ? 1 : 0)
                }
                .fixedSize() 
                
                ZStack {
                    if isSelected {
                        Rectangle()
                            .fill(sakuraPink)
                            .frame(height: 3)
                            .matchedGeometryEffect(id: "indicator", in: animationNamespace)
                    } else {
                        Rectangle()
                            .fill(.clear)
                            .frame(height: 3)
                    }
                }
            }
            .contentShape(Rectangle())
        }
    }
}

#Preview {
    FanclubTabView()
}
