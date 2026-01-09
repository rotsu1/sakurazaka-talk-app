//
//  TalkTabView.swift
//  sakurazaka-talk-app
//
//  Created by 乙津　龍　 on 3/1/2026.
//

import SwiftUI
import SwiftData
import Kingfisher




struct TalkTabView: View {
    @Query(sort: \Member.joinOrder) private var members: [Member]
    @Environment(\.modelContext) private var modelContext

    @StateObject private var subService = SubscriptionService()

    @State private var showModal = false
    @State private var selectedMember: Member? = nil

    private var groupedMembers: [(generation: String, list: [Member])] {
        let grouped = Dictionary(grouping: members) { member in
            if subService.isSubscribed(to: member) {
                return 0
            } else {
                return member.joinOrder
            }
        }
        
        let sortedKeys = grouped.keys.sorted()
        
        return sortedKeys.map { key in
            let label = key == 0 ? "オンライン" : "\(key)期生"
            let sortedList = (grouped[key] ?? []).sorted(by: { $0.name < $1.name })
            return (generation: label, list: sortedList)
        }
    }

    let columns = [
        GridItem(.flexible()),
        GridItem(.flexible()),
        GridItem(.flexible()),
    ]

    var body: some View {
        ZStack {
            VStack {
                HeaderView(title: "トーク", icons: true, isBlog: false, isSubpage: false)

                ScrollView {
                    LazyVGrid(columns: columns, spacing: 16) {
                        ForEach(groupedMembers, id: \.generation) { memberGroup in
                            Section(
                                header: GenerationHeader(title: memberGroup.generation),
                                footer: Group {
                                    if memberGroup.generation != groupedMembers.last?.generation {
                                        Rectangle()
                                            .fill(Color.gray.opacity(0.2))
                                            .frame(height: 2)
                                            .padding(.horizontal, 20)
                                    }
                                }
                                ) {
                                ForEach(memberGroup.list) { member in
                                    VStack {
                                        if memberGroup.generation == "オンライン" || subService.isSubscribed(to: member) {
                                            NavigationLink(destination: TalkView(member: member)) {
                                                MemberSection(member: member)
                                            }
                                        } else {
                                            Button {
                                                showModal = true
                                                selectedMember = member
                                            } label: {
                                                MemberSection(member: member)
                                            }
                                        }
                                    }
                                    .foregroundColor(sakuraPink)
                                    .padding(.bottom, 16)
                                    .id("\(member.id)-\(member.subscription?.status ?? "nil")")
                                }
                            }
                        }
                    }
                }
                .scrollIndicators(.hidden)
            }
            .padding(.bottom, 110)

            if showModal {
                SubscriptionModal(isPresented: $showModal, member: selectedMember, service: subService)
            }
        }
        .task {
            let service = MemberService(modelContext: modelContext)
            do {
                try await service.syncMembers()
            } catch {
                print("❌ Sync failed: \(error)") // This will print the actual reason
            }
        }

    }
}

struct MemberSection: View {
    var member: Member

    var body: some View {
        VStack {
            KFImage(URL(string: member.avatarUrl))
                .placeholder { 
                    Circle().fill(Color.gray.opacity(0.3))
                }
                .aspectRatio(contentMode: .fill)
                .frame(width: 96, height: 96)
                .clipShape(Circle())
            
            Text(member.name)
                .font(.system(size: 14))
                .foregroundColor(sakuraPink)
        }
    }
}

// A helper view to style generation headers
struct GenerationHeader: View {
    let title: String
    
    var body: some View {
        HStack {
            Text(title)
                .font(.headline)
                .padding(.horizontal, 16)
                .padding(.vertical, 16)
                .background(Color.white.opacity(0.9))
                .foregroundColor(sakuraPink) 
        }
        .frame(maxWidth: .infinity)
        .background(Color.white) 
    }
}

#Preview {
    TalkTabView()
}