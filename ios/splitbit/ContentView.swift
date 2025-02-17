import SwiftUI

struct ContentView: View {

    
    var body: some View {
        
        TabView {
            // First Tab
            NavigationView {
                GroupListView()
                    .navigationTitle("Your Groups")
            }
            .tabItem {
                Image(systemName: "house")
                Text("Groups")
            }


            // Second Tab
            
            
            
            NavigationView {
                Addpay()
                    .navigationTitle("Add New Item")
            }
            .tabItem {
                Image(systemName: "plus")
                Text("Add")
            }
            // third tab
            
            NavigationView {
                SettingsView()
                    .navigationTitle("Settings")
            }
            .tabItem {
                Image(systemName: "gear")
                Text("Settings")
            }
        }
    }
}







struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
