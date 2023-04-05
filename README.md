# AEM-Detector

The tool is written in the Go programming language, designed to discover Adobe Experience Manager (AEM) content management system (CMS) websites. It takes a list of URLs as input and uses various techniques to determine whether the website is using Adobe Experience Manager.

The tool is able to identify AEM websites by analyzing the patterns and URLs that are unique to AEM websites. The tool is designed to be lightweight, fast, and easy to use, with a simple command-line interface that allows users to quickly scan a large number of URLs for AEM websites.

# Installation
```
go install github.com/akshaysharma016/aem-detector@latest

```

# How to Run
```
cat live-domains | aem_detector
```

# Thanks

Special thanks to “Mikhail Egorov” (https://twitter.com/0ang3el) for his wonderful research on AEM CMS. This tool is designed using his research (https://github.com/0ang3el/aem-hacker) on AEM CMS. It's always great to see researchers like him contributing to the open-source community and sharing their work with others. 

If you have found the tool to be useful, you can show your appreciation by leaving a positive comment or review, sharing the tool with others, or even contributing. These actions can help motivate researchers to continue their work and inspire others to contribute to the community as well.
