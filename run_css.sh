
echo
echo "Starting TailwindCSS compiler ..."

npx tailwindcss@3.4.17 -c ./web_src/tailwind.config.js -i ./web_src/main.css -o ./web/main.css --watch
