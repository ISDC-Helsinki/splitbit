#!/bin/sh

if ! command -v sqlbs >/dev/null 2>&1; then
        echo -e "No sqlbs command found. Visit: \033[1mhttps://github.com/Kuchteq/sqlbs \033[mand get the newest release place it in one of the PATH directories and make it executable or clone the repository and run go build on it." 
        echo "Alternatively, make the script attempt to download the newest release [y/n]"
        read ans
        if [ "$ans" = y ]; then
                curl -LO https://github.com/Kuchteq/sqlbs/releases/latest/download/sqlbs && chmod u+x sqlbs
                if echo "$PATH" | grep -q "$HOME/.local/bin"; then 
                        mv sqlbs "$HOME/.local/bin" 
                else
                        echo "No local bin executable folder, placing in /usr/local/bin this will require sudo privilages"
                        sudo mv sqlbs "/usr/local/bin" 
                fi
        fi
        exit 1
fi

if ! [ -f "data.db" ]; then
        sqlite3 data.db < schema.sql
fi

sqlbs ./schema.sql | sqlite3 data.db
echo "If you see any errors related to UNIQUE constraint those can be safely ignored. Sqlbs is still in development but those shouldn't affect much."
echo -e "\033[1mSample data successfully inserted!\033[m"
