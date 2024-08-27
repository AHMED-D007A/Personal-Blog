# Personal Blog
- Simple personal blog where I can write and publish articles.
- This project is from the project based road map, you can find it [here](https://roadmap.sh/backend/projects).
- You can read more about the project form [here](https://roadmap.sh/projects/personal-blog).
- App pages are renderd using templ which is a go package for Createing components that render fragments of HTML and compose them to create screens, pages, documents, or apps.
- For more about templ you can visit this website [here](https://templ.guide).

### About
- The blog will have two sections: a guest section and an admin section.
  - Guest Section — A list of pages that can be accessed by anyone:
    - Home Page: This page will display the list of articles published on the blog.
    - Article Page: This page will display the content of the article along with the date of publication.
  - Admin Section — are the pages that only you can access to publish, edit, or delete articles.
    - Dashboard: This page display the list of articles published on the blog along with the option to add a new article, edit an existing article, or delete an article.
    - Add Article Page: This page contain a form to add a new article. The form will have fields like title, content, and date of publication.
    - Edit Article Page: This page contain a form to edit an existing article. The form will have fields like title, content, and date of publication.

### Usage
- Download repository's files, navigate to the folder and then run ```make```
- After that open your browser and navigate to this link ```http://localhost:4000```
