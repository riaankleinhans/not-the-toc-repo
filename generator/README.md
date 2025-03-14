# TAG and TOC Subproject README Generator

This folder contains a Go program (`reamme_app.go`) that automatically generates README files for Technical Advisory Groups (TAGs) and TOC Subprojects based on data defined in `tags.yaml`.

The authoritative source for TAG and TOC Subproject information is the `tags.yaml` file in the project root. All updates to TAG and TOC Subproject details must be made in this file.

### `tags.yaml` Structure

The structure of the `tags.yaml` file is as follows:

```yaml
tags:
  - dir: "tag-example"
    name: "Example TAG"
    mission_statement: "This is an example TAG."
    leadership:
      chairs:
        - name: "John Doe"
          github: "johndoe"
    meetings:
      - description: "TAG Meeting"
        tag_calendar: https://zoom-lfx.platform.linuxfoundation.org/meetings/cncf?view=week
        recordings_url: https://www.youtube.com/playlist?foo
    contact:
      slack: "#example-tag"
      mailing_list: "[email address removed]"
    tag_subprojects:
      - name: "Example Subproject"
        contact:
          slack: "#example-subproject"
          mailing_list: "[email address removed]"
toc_subprojects:
  - dir: "toc-subproject-example"
    name: "Example TOC Subproject"
    mission_statement: "This is an example TOC Subproject."
```

### README Generation
The program uses a Go template (generator/tag_readme.tmpl) to generate the README files for each TAG.

### How to Update READMEs

* Modify tags.yaml: To update a README file, make the necessary changes to the corresponding TAG or TOC Subproject entry in the tags.yaml file.
* Submit a Pull Request (PR): Create a pull request with your changes to the tags.yaml file.
* Automation: The automation system will automatically run the reamme_app.go script.
* Updated READMEs in PR: The automation will then commit and push the updated README files directly to your pull request.
* Review and Merge: Review the changes made by the automation, and if everything looks correct, merge your pull request.