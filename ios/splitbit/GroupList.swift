//
//  GroupList.swift
//  splitbit
//
//  Created by Anastasia on 28/10/2024.
//

import SwiftUI

struct GroupItem: Identifiable {
    let id = UUID()
    let name: String
}

struct GroupListView: View {
    let groups = [
        GroupItem(name: "Home group"),
        GroupItem(name: "Group 2"),
        GroupItem(name: "Group 3")
    ]

    var body: some View {
        NavigationView {
            List(groups) { group in
                NavigationLink(destination: destinationView(for: group)) {
                    HStack{
                        Image(systemName: "rays")
                            .font(.system(size: 30))
                        Text(group.name)
                            .font(.system(size: 25, weight: .bold, design: .rounded))
                            .foregroundColor(Color("Text colour"))
                        
                        
                    }
                }
                .padding(.vertical, 15)
                .listRowSeparator(.hidden)
            }
            .navigationTitle("Groups")
        }
    }
    
    
    

    
    @ViewBuilder
    private func destinationView(for group: GroupItem) -> some View {
        if group.name == "Home group" {
            Homegroup()
        } else {
            HomegroupView(group: group)
        }
    }
}

struct HomegroupView: View {
    let group: GroupItem

    var body: some View {
        VStack {
            Text("Welcome to \(group.name)")
                .font(.largeTitle)
                .padding()
        }
        .navigationTitle(group.name)
    }
}
