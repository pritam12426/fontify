import json

# wget "https://formulae.brew.sh/api/cask.json" -P "/Users/pritam/Library/Caches/fontify" -O "homeBrew-Fontify.json"

with open("/Users/pritam/Library/Caches/Fontify/homeBrew-Fontify.json") as f:
	data: list = json.load(f)

newData = []
raw_data: str = ""

ignore: int = 0
for i in (data):
	artifacts = i["artifacts"][0].keys()

	for j in artifacts:
		if j == "font":
			if i["url"].endswith(".git") or i["url"].endswith(".exe"):
				ignore += 1
				continue
			else:
				# print(f"{', '.join(i['name'])} :: {i['url']} ")
				raw_data += f"{', '.join(i['name'])} :: {i['url']}"
				newData.append(i)
		else:
				ignore += 1



print(f"\nIgnore {ignore}")
print(f"Total fonts {len(newData)}")
h = json.dumps(newData, indent="\t")

with open("temp-new-font-data.json", "w") as f:
	f.write(h)

with open("font_list.txt", "w") as f:
	f.write(raw_data)
