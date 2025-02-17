//
//  splitbitApp.swift
//  splitbit
//
//  Created by Anastasia on 21/10/2024.
//

import SwiftUI

@main
struct splitbitApp: App {
    var body: some Scene {
        WindowGroup {
            ContentView()
                .environment(\.font, Font.custom("HelveticaNeue-Bold", size: 16))
        }
    }
}

struct splitbitApp_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
            .environment(\.font, Font.custom("HelveticaNeue-Bold", size: 16))
    }
}
