# Ryan Parman • [jobs@ryanparman.com](mailto:jobs@ryanparman.com)

**Cloud-native engineering leader with a focus on reliability, scalability, and security for the modern web.**

**Links:** [GitHub (personal)] • [GitHub (side project)] • [LinkedIn] • [Stack Overflow] • [Role-targeted résumés](https://github.com/skyzyx/resume/blob/master/resumes/#readme) \
**Format:** [Web](https://github.com/skyzyx/resume/blob/master/resumes/ryanparman-software-eng-devtools.md) • [PDF](https://github.com/skyzyx/resume/raw/master/resumes/ryanparman-software-eng-devtools.pdf) • [Word](https://github.com/skyzyx/resume/raw/master/resumes/ryanparman-software-eng-devtools.docx) • [OpenDocument](https://github.com/skyzyx/resume/raw/master/resumes/ryanparman-software-eng-devtools.odt)

## Summary

Ryan Parman is a cloud-native engineering leader, who specializes in technical leadership, software development, site reliability engineering, and cybersecurity for the modern web. Excels at listening, adapting, and driving continuous improvement.

Small business owner, two-time startup founder, and creator of two open-source projects with millions of users each. Ryan has a proven track record of building high-quality software, delivering impactful solutions, and elevating team performance.

## Work Experience

Older roles are truncated for brevity. If interested, details can be found [on GitHub](https://github.com/skyzyx/resume/blob/master/ryanparman-previously.md).

### [McGraw Hill] — Remote (since COVID), previously Seattle, WA

> McGraw Hill is a _learning science_ company which produces textbooks, digital learning tools, and adaptive technology to enhance learning. It is one of the “big three” educational publishers in the U.S, and was acquired by Platinum Equity 2021.

#### Principal Engineer, Cloud Center of Excellence (January 2024—October 2024)

* Joined a team whose mission was to provide guidance and support in the cloud journey of the entire organization.

* Proposed best practices, guardrails, and security measures to ensure a secure and efficient cloud environment.

* Identified opportunities to extend the security measures and guardrails devised for AWS to other cloud platforms.

#### Principal Cloud and Platform Engineer (June 2020—January 2024)

* As every school in America transitioned to online learning during the COVID-19 lockdowns, I was the technical/development lead on the team who supported all SRE and product engineering teams, working on core platforms and services.

* Managed the Base [AMI] program. Leveraged insights from [CIS], security patching, and internal needs to develop a unified build pipeline integrating best practices.

* Conducted comprehensive scans of [Route 53][Amazon Route 53] to obtain a mapping of the company’s thousands of active websites. Prioritized identifying and remediating misconfigurations, rotating certificates, and increasing visibility.

* Spearheaded the [Artifactory] Rebuild project. Ran the project from inception to completion, including the majority of development. Directed effort across ~80 teams and ~300 services to complete the project.

* Led dozens of smaller projects, offered guidance to engineers on best practices, and documented knowledge.

#### Engineering Manager, Site Reliability (October 2018—June 2020)

* Led the [_Site Reliability Engineering_][SRE] (SRE) team in addressing macro-oriented problems affecting engineering teams, empowering greater self-service.

* Established a process for maintaining reusable [Terraform] modules which teams leveraged to compose infrastructure with minimal effort.

* Customized the [Amazon Linux] AMIs to comply with Level-2 [CIS] Guidelines for both Amazon Linux and [Docker]. Liaised with cybersecurity, operations, and business units to ensure compliance.

* Invented custom security and operational tooling to understand the current posture of ~200 AWS accounts where off-the-shelf tools did not meet the needs of the organization.

* Reduced the time to deploy a new service from dozens of weeks to a single meeting by implementing a _Monitoring-as-Code_ methodology, and defining broad-use [Service Level Objectives][SLO] (SLOs).

<div class="page-break"></div>

#### Staff Software Engineer (October 2016—October 2018)

* Led the development of Tier-1 services within the educational content authoring pipeline, leveraging technologies such as [REST], [GraphQL], API design, [Amazon ECS] (similar to [Kubernetes]), [Docker], [Terraform], [ePubs][EPUB], and security best practices.

* Led the development of the authoring component of the [SmartBook 2.0 product][SB2], and the internal system which indexes authored content, builds ePubs, and encodes images/video for the ePub CDN using [ffmpeg].

* Established the technical direction of these projects, promoted adoption across the organization, published comprehensive documentation, and offered ongoing integration guidance.

* Accelerated the adoption of CI/CD, rapid deployment practices, and Docker containers, shortening the feedback loop for developers and increasing the reliability of deployments.

### [WePay] — Redwood City, CA

> WePay is an online payment service provider which provides “payments for platforms”, where examples of platforms are GoFundMe, Care.com, and Xbox. It was acquired by JPMorgan Chase in October 2017.

#### DevOps Engineer (April 2015—September 2016)

* Led a cross-company initiative to upgrade the monolithic application from [PHP] 5.4 to PHP 5.6 (the latest at the time). Facilitated cross-team collaboration among all major engineering teams and QA departments to achieve results.

* Initiated a program to automate the creation of base server images for cloud servers. This allowed new servers to boot and begin serving traffic ~75% faster.

* Invested in monitoring and alerting systems to prevent customer-facing issues.

#### Senior API Engineer (April 2014—April 2015)

* Led the company’s [HackerOne](https://www.hackerone.com) program, coordinating across teams to address security issues.

* Built a development environment for engineering teams. Reduced new engineer onboarding time from 2 weeks to 1 day.

* Expanded WePay’s payment security offerings by designing MFA-as-a-Service (U.S. patent filing [US15042104]).

### [Amazon Web Services] — Seattle, WA

> Amazon Web Services provides on-demand cloud computing platforms and APIs to individuals, companies, and governments, on a metered, pay-as-you-go basis.

#### Web Development Engineer II (March 2010—April 2014)

* Created the [AWS SDK for PHP], enabled AWS to reach the largest developer group — [PHP].

* Initiated the creation of [AWS SDK for PHP] v2 to address changes in the PHP language and growth of AWS services.

* Led one of the first teams to provide reusable UI building blocks for the [AWS Management Console], by collaborating directly with the AWS Design team.

* Invested in increased transparency, better communication, and improved tooling for developers. [[Examples](https://github.com/skyzyx/resume/blob/master/amazon-specifics.md)]

## Projects

Proof that I can code, call APIs, interact with SDKs, and build user-facing software. I have live-coding anxiety, so live-coding interviews will always present me at my worst, not my best.

* **DevSec Tools:** Building a [website](https://github.com/northwood-labs/devsec-ui), [CLI tool, and Go library](https://github.com/northwood-labs/devsec-tools) for helping developers identify potential web security configuration issues (in-progress).

* **Custom Linux Packages:** Building a [repository of custom Linux packages](https://github.com/northwood-labs/package-building/wiki) (in-progress).

* **CSP Evaluator:** Building a [parser and evaluator for Content Security Policy (CSP) directives](https://github.com/northwood-labs/csp-parser) in Go (in-progress).

* **Terraform Provider:** Built a [custom provider](https://github.com/northwood-labs/terraform-provider-corefunc) which provides a set of utility functions for use in Terraform/OpenTofu.

* **Multi-Platform Docker:** Built a [downloader for GitHub release assets](https://github.com/northwood-labs/download-asset) which simplifies building multi-platform images.

* **AWS Organization Security:** Built a [library + CLI tool](https://github.com/northwood-labs/assume-spoke-role) which simplifies the AWS pattern for multi-account organizations which they call “hub and spoke.”

* **AWS Session Manager:** The terminal is the right tool for shell sessions. Built a [TUI](https://en.wikipedia.org/wiki/Text-based_user_interface) for [simplifying connections to SSM-enabled EC2 instances](https://github.com/northwood-labs/ssm-shell) using your Terminal.

* **Configuration for `tflint`:** Built a [tool for generating up-to-date configurations for AWS/GCP/Azure](https://github.com/northwood-labs/tflint-config-generator) for use with [tflint](https://github.com/terraform-linters/tflint).

## Recommendations

See a [selective list of recommendations](https://github.com/skyzyx/resume/blob/master/selected-recommendations.md) from co-workers and peers.

## Groups and Accomplishments

* U.S. patent filing, [“System and Methods for User Authentication across Multiple Domains”][US15042104] (US15042104) (2016)
* U.S. patent filing, [“Hive-based Peer-to-Peer Network”][US8103870B2] (US8103870B2) (2007)
* Voting representative for AWS, [PHP Framework Interoperability Group][PHP-FIG] (2012–2013)

## Skills

This list is not exhaustive, but these are software and skills I leveraged in the roles above which are most relevant to Software Engineering and DevTools roles.

API design, API versioning, CLI tools, [Bash], [CircleCI], [Docker], [GitHub Actions], [Git], [Go], [GraphQL], [JWT], [NFS], [PHP], [Python], [REST], [Redis], [Subversion], [Vagrant], [WordPress], [XSLT], [ffmpeg], [twelve-factor applications], automation, building platforms, code generation, defensive cybersecurity, multi-platform development, performance, scalability, software library design, software testing, technical documentation.

## Education

_Silicon Valley College_ (now [Carrington College]), San Jose, CA. Bachelor of Arts, _Design and Visualization_

[Alpine Linux]: https://alpinelinux.org
[Amazon ACM]: https://aws.amazon.com/certificate-manager/
[Amazon CloudFront]: https://aws.amazon.com/cloudfront/
[Amazon EC2]: https://aws.amazon.com/ec2/
[Amazon ECS]: https://aws.amazon.com/ecs/
[Amazon IAM]: https://aws.amazon.com/iam/
[Amazon Linux]: https://aws.amazon.com/linux/
[Amazon Route 53]: https://aws.amazon.com/route53/
[Amazon S3]: https://aws.amazon.com/s3/
[Amazon Web Services]: https://www.crunchbase.com/organization/amazon-web-services
[AMI]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/AMIs.html
[Ansible]: https://www.redhat.com/en/technologies/management/ansible
[ARM64]: https://aws.amazon.com/ec2/graviton/
[Artifactory]: https://jfrog.com/artifactory/
[AWS CloudFormation]: https://aws.amazon.com/cloudformation/
[AWS Control Tower]: https://aws.amazon.com/controltower/
[AWS Elastic Beanstalk]: http://aws.amazon.com/elasticbeanstalk/
[AWS Identity Center]: https://aws.amazon.com/iam/identity-center/
[AWS Lambda]: https://aws.amazon.com/lambda/
[AWS Management Console]: https://console.aws.amazon.com
[AWS Professional Services]: https://aws.amazon.com/professional-services/
[AWS RDS Aurora]: https://aws.amazon.com/rds/aurora/
[AWS SDK for PHP]: https://aws.amazon.com/sdk-for-php/
[AWS SDKs]: https://aws.amazon.com/developer/tools/
[AWS Secrets Manager]: https://aws.amazon.com/secrets-manager/
[AWS Well-Architected]: https://aws.amazon.com/architecture/well-architected/
[AWS]: https://aws.amazon.com
[Bash]: https://www.gnu.org/software/bash/
[Carrington College]: http://carrington.edu/schools/san-jose-california/
[CentOS]: https://en.wikipedia.org/wiki/CentOS
[CircleCI]: https://circleci.com/enterprise/
[CIS]: https://www.cisecurity.org
[Confluence]: https://www.atlassian.com/software/confluence
[Docker]: https://www.docker.com
[EC2 Image Builder]: https://aws.amazon.com/image-builder/
[EPUB]: https://www.w3.org/publishing/epub3/
[ffmpeg]: https://ffmpeg.org
[GCP]: https://cloud.google.com
[Git]: https://git-scm.com
[GitHub (personal)]: https://github.com/skyzyx
[GitHub (side project)]: https://github.com/northwood-labs/
[GitHub Actions]: https://github.com/features/actions
[GitHub Enterprise]: https://github.com/enterprise
[Go]: https://go.dev
[GraphQL]: https://graphql.org
[Jenkins]: https://www.jenkins.io
[Jira]: https://www.atlassian.com/software/jira
[JWT]: https://jwt.io
[kubectl]: https://github.com/kubernetes/kubectl
[Kubernetes]: https://kubernetes.io
[LinkedIn]: https://www.linkedin.com/in/rparman/
[McGraw Hill]: https://www.crunchbase.com/organization/mcgraw-hill-education
[NFS]: https://en.wikipedia.org/wiki/Network_File_System
[Nginx]: https://www.nginx.com
[Northwood Labs]: https://www.crunchbase.com/organization/northwood-labs
[OpenTofu]: https://opentofu.org
[Packer]: https://packer.io
[PHP-FIG]: http://www.php-fig.org
[PHP]: https://www.php.net
[Python]: https://www.python.org
[Redis]: https://redis.io
[REST]: https://martinfowler.com/articles/richardsonMaturityModel.html
[SB2]: https://www.mheducation.com/news-media/press-releases/mcgraw-hill-connect-unveils-smartbook.html
[SLO]: https://sre.google/sre-book/service-level-objectives/
[SRE]: https://sre.google/in-conversation/
[Stack Overflow]: https://stackoverflow.com/users/228514/ryan-parman
[Subversion]: https://subversion.apache.org
[Terraform]: https://www.terraform.io
[twelve-factor applications]: https://12factor.net
[Ubuntu]: https://ubuntu.com
[US15042104]: https://patents.google.com/patent/US20160241536A1/en?inventor=Ryan+Parman
[US8103870B2]: https://patents.google.com/patent/US8103870B2/en?inventor=Ryan+Parman
[Vagrant]: https://www.vagrantup.com
[WarpShare]: https://www.crunchbase.com/organization/warpshare
[WePay]: https://www.crunchbase.com/organization/wepay
[WordPress]: https://wordpress.org
[XSLT]: https://developer.mozilla.org/en-US/docs/Web/XSLT
