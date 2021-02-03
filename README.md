# Go-Service

This repository holds basic template to write golang microservice.

### How to use this simple template

This template is designed to be used when new github repository is created. Follow next steps to get ready for code your project:

1. Clone this repo:

   git clone https://github.com/boletia/go-service.git

2. Rename cloned dir to match name with your new project:

   `mv go-service myNewProject`

3. Enter into your renamed project directory:

   `cd myNewProject`

4. Change **APPNAME** variable value from **go-service** to **YOUR project name** into Makefile file:

   `perl -pi -e 's/go-service/myNewProject/' Makefile`

5. Rename the directory under cmd directory replacing **go-service** name for **YOUR project name**:

   `mv cmd/go-service cmd/myNewProject`

6. Make sure tests and lint make targets work properly:

   `make test && make lint && make build`

7. Delete **.git** dir and init git again:

   `rm -rf .git; git init`

8. Clean README.md file:

   `echo "firts commit" > README.md`

9. Add all files and make your first commit:

   `git add .; git commit -m 'initial-commit'`

10. Add the corresponding remote to your new git local repo:

    `git remote add origin https://github.com/boletia/<YOUR_NEW_PROJECT>.git`

11. push your first commit to remote repo:

    `git push origin master`



### Enabling required status check

In order to improve our code we could enable checks to ensure the new code pass all unit tests written for it, and there is not complains about linting process, we can use the same makefile tarjets described above, `make test && make lint && make build` to deny merge process when our code doesn't pass unit-test, has linting problems or the binary file is not able to be built from source, to do so follow next steps:

Protect master branch:

- Go to the main page of the repository and click over **Settings**.
- In the left menu, click **Branches**.
- Next to "Branch protection rules", click **Add rule.**
- Under "Branch name pattern", type the branch name **master**.
- Check **Require pull request reviews before merging** option.
- Check **Require status checks to pass before merging** option.
- Check **test-lint-build** status check which appears into **Status checks found in the last week for this repository** box.
- Click over **Create** button which apperas at bottom from the page.
- Click over **Save changes** button which apperas at bottom from the page.
- In your local machine create a new branch to work with for you next project feature.



Next time you create a poll request github will launch the action created and the code will be validated, if unit tests don't pass or if linting process complains about the code or if a build go process are not able to generate go binary, merge process will be denied until you can fix the code.



### Makefile tarjets

1. **build**: build your project creating binary output file into **build/** dir.
2. **run**: build & run built binary file.
3. **test**: read and execute all **\*_test.go** files, run unit-tests.
4. **test-report**: read output report file and show unit-tests coverage.
5. **lint**: execute [golangci-lint](https://github.com/golangci/golangci-lint) linter over the code to check linting issues.
6. **clean**: remove go project binary file if it exists.



### TODO:

- Document directory layout.

  

  