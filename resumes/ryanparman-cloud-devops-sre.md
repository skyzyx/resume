# Ryan Parman • [jobs@ryanparman.com](mailto:jobs@ryanparman.com)

**Cloud Engineering Leader • Innovator • Problem Solver; looking for roles in technical leadership and engineering management.**

**Links:** [GitHub (personal)] • [GitHub (side project)] • [LinkedIn] • [Stack Overflow] • [Role-targeted résumés](https://github.com/skyzyx/resume/blob/master/resumes/#readme) \
**Format:** [Web](https://github.com/skyzyx/resume/blob/master/resumes/ryanparman-cloud-devops-sre.md) • [PDF](https://github.com/skyzyx/resume/raw/master/resumes/ryanparman-cloud-devops-sre.pdf) • [Word](https://github.com/skyzyx/resume/raw/master/resumes/ryanparman-cloud-devops-sre.docx)

## Summary

Cloud engineering leader with a diverse background spanning design, development, security, and innovation. Proven expertise in building scalable infrastructure, driving efficiency, and enhancing user experience. Adept at leading teams, streamlining complex processes, and fostering knowledge-sharing cultures. Passionate about solving real-world problems through technology, security, and strategic thinking.

### Key Skills

<table width="100%">
<tbody>
<tr>
<td>

* Cloud Engineering and Infrastructure
* Security and Compliance

</td>
<td>

* Technical Leadership and Team Building
* Documentation and Knowledge Sharing

</td>
<td>

* Cost Optimization and Strategic Planning
* Developer Productivity and Tooling

</td>
</tr>
</tbody>
</table>

## Work Experience

### [Stripe] — Remote

> Stripe provides a fully integrated suite of financial and payments products.

#### Senior Technical Program Manager, Databases (March 2025—Present)

* Managed weekly, monthly, and quarterly budgets for the Databases organization, on the order of millions of dollars per month, keeping the budget "in the green" and diving deeply into erroneus spending.
* Managed multiple internal migration projects, including tracking progress and relevant metrics, working across several teams to achieve the goal, and reducing how much work our infrastructure teams pushed out to product teams to perform by focusing on automation solutions and understanding the ROI/impact of the ask.
* Managed and performed the work to migrate all internal-user facing documentation from multiple systems/locations into a singular system. Performed technical editing, and leveraged a [Diátaxis](https://diataxis.fr)-like approach to content organization.

### [McGraw Hill] — Remote (since COVID), previously Seattle, WA

> McGraw Hill is a _learning science_ company which produces textbooks, digital learning tools, and adaptive technology to enhance learning. It is one of the “big three” educational publishers in the U.S.

#### Principal Engineer, Cloud Architecture (January 2024—October 2024)

* Joined a team whose mission was to provide guidance and support in the cloud journey of the entire organization.
* Started development on v2 of a project which scanned AWS accounts for misconfigurations and vulnerabilities. Goal was to reduce an [AWS Well-Architected] review from 2 weeks to 2 hours with automated scans, enabling more reviews annually (10 → 100).
* Managed the migration from [CentOS] to [Amazon Linux] before the CentOS end-of-life date, which elevated our security posture.
* Proposed best practices, guardrails, and security measures to ensure a secure and efficient cloud environment.
* Identified opportunities to extend the security measures and guardrails devised for AWS to other cloud platforms (Microsoft Azure, Oracle Cloud).

#### Principal Cloud and Platform Engineer (June 2020—January 2024)

* Led the team who supported all SRE and product engineering teams, scaling core platforms and services as every school in America transitioned to online learning during the COVID-19 lockdowns.
* Partnered with Enterprise Architecture and [AWS Professional Services] to deploy [Control Tower][AWS Control Tower] and [Identity Center][AWS Identity Center], resulting in lowered costs and improved control over security guardrails.
* Managed the Base [AMI] program (server disk images). Leveraged insights from [CIS], security patching, and internal needs to develop a unified build pipeline integrating best practices. Reduced time-to-boot, and eliminated engineering toil (1 → 10).
* Conducted comprehensive scans of [Route 53][Amazon Route 53] to obtain a mapping of the company’s thousands of active websites. Prioritized identifying and remediating misconfigurations, rotating certificates, and increasing visibility.
* Grew and ran a project which evaluated AWS accounts for high-priority misconfigurations and vulnerabilities. Included a high-level score (friendly competition), explanations of the issues (security education), and instructions for fixing (driving forward). Became a trusted tool across the organization (1 → 10).
* Implemented the custom Linux runtime environment used by self-hosted [GitHub Actions] runners.
* Spearheaded the [Artifactory] Rebuild project. Ran the project from inception to completion, including the majority of development. Directed effort across ~80 teams and ~300 services to complete the project.
* Enabled continuous token and password rotation for engineering teams by designing and deploying a _Token Vending Machine_, improving security.
* Adapted the _Monitoring-as-Code_ tooling/methodology to abstract-away the underlying vendor, streamlining a vendor migration ([New Relic], [Datadog]) (1 → 10).
* Resolved all technology blockers preventing migration lower-cost [ARM64] CPUs, opening the door for ~$450k/year in cost savings.
* Led dozens of smaller projects, offered guidance to engineers on best practices, and documented knowledge.

#### Site Reliability Engineering Manager (October 2018—June 2020)

* Led the [_Site Reliability Engineering_][SRE] (SRE) team in addressing macro problems affecting engineering, empowering self-service.
* Established a process for maintaining reusable [Terraform] modules which teams leveraged to compose infrastructure.
* Customized the [Amazon Linux] AMIs to comply with Level-2 [CIS] Guidelines for both Amazon Linux and [Docker]. Liaised with cybersecurity, operations, and business units to ensure compliance (0 → 1).
* Invented custom security and operational tooling to understand the current posture of AWS accounts where off-the-shelf tools did not meet the needs of the organization (0 → 1).
* Reduced the time to deploy a new service from dozens of weeks to a single meeting by implementing a _Monitoring-as-Code_ methodology, and defining broad-use [Service Level Objectives][SLO] (SLOs) ([New Relic], [Datadog]) (0 → 1).

#### Staff Software Engineer (October 2016—October 2018)

* Led the development of Tier-1 services within the educational content authoring pipeline, leveraging technologies such as [REST], [GraphQL], API design, [Amazon ECS] (similar to [Kubernetes]), [Docker], [Terraform], [ePubs][EPUB], and security best practices.
* Led the development of the authoring component of the [SmartBook 2.0 product][SB2], and the internal system which indexes authored content, builds ePubs, and encodes images/video for the ePub CDN using [ffmpeg].
* Established the technical direction of these projects, promoted adoption across the organization, published comprehensive documentation, and offered ongoing integration guidance.
* Accelerated the adoption of CI/CD, rapid deployment practices, and Docker containers, shortening the feedback loop for developers and increasing the reliability of deployments.
* Served as a core resource in adopting _Infrastructure-as-Code_ (IaC) tools such as [Terraform] and [Packer].

### [WePay] — Redwood City, CA

> WePay is an online payment service provider which provides “payments for platforms”, where examples of platforms are GoFundMe, Care.com, and Xbox. JPMorgan Chase acquired WePay in October 2017.

#### Senior DevOps Engineer (April 2015—September 2016)

* Led a cross-company initiative to upgrade the monolithic application from [PHP] 5.4 to PHP 5.6 (the latest at the time). Facilitated cross-team collaboration among all major engineering teams and QA departments to achieve results.
* Initiated a program to automate the creation of base server images for cloud servers. This allowed new servers to boot and begin serving traffic ~75% faster.
* Invested in monitoring and alerting systems to prevent customer-facing issues ([New Relic], [Grafana]).
* Explored _configuration-as-code_ for cloud infrastructure in [Google Cloud Platform][GCP] to improve reliability and efficiency.

#### Senior API Engineer (April 2014—April 2015)

* Led the company’s [HackerOne](https://www.hackerone.com) security program, coordinating across teams to address security issues.
* Built a local development environment for engineering teams using [Vagrant]. Eliminated "works on my machine", and reduced new engineer onboarding time from 2 weeks to 1 day (measured by when a new employee could make their first commit).
* Expanded WePay’s payment security offerings by designing MFA-as-a-Service (U.S. patent filing [US15042104]).

<div class="page-break"></div>

### Older roles, side projects

See “[Previous experience, side projects](https://github.com/skyzyx/resume/blob/master/ryanparman-previously.md)” for additional details.

* [Northwood Labs](https://github.com/northwood-labs) — Owner (January 2024—Present)
* PCR Publishing (Side-Project) — Editor, Typesetter, Publisher, Book Producer (April 2021–April 2022)
* Perimeter of Wisdom, LLC (defunct) — Co-Owner, CTO, Producer (February 2015—2018)
* [Amazon Web Services] — AWS SDK Developer (March 2010—April 2014)
* Rearden Commerce (now [Deem](https://www.crunchbase.com/organization/deem)) — Senior User Experience Developer (July 2008—March 2010)
* [WarpShare](https://www.crunchbase.com/organization/warpshare) (defunct) — Co-Founder and Chief Information Officer (September 2006—March 2010)
* [Yahoo!](https://www.crunchbase.com/organization/yahoo) — Front-end Developer (Contract), Yahoo! Messenger (November 2007—January 2008)
* [Stryker](https://www.crunchbase.com/organization/stryker) — User Interface Developer (Contract) (May 2005—September 2006)
* [Digital Impact](https://www.crunchbase.com/organization/digital-impact-2) (now part of [Axciom](https://www.crunchbase.com/organization/acxiom-digital-inc)) — Production Specialist (March 2004—April 2005)

## Projects

Proof that I can code, call APIs, interact with SDKs, and build user-facing software. I have live-coding anxiety, so live-coding interviews will always present me at my worst, not my best.

* **DevSec Tools:** Building a [website](https://github.com/northwood-labs/devsec-ui), [CLI tool, and Go library](https://github.com/northwood-labs/devsec-tools) for identifying potential security configuration issues (in-progress).
* **Custom Linux Packages:** Building a [repository of custom Linux packages](https://github.com/northwood-labs/package-building/wiki) (in-progress).
* **Terraform Provider:** Built a [custom provider](https://github.com/northwood-labs/terraform-provider-corefunc) which provides a set of utility functions for use in Terraform/OpenTofu.
* **Multi-Platform Docker:** Built a [downloader for GitHub release assets](https://github.com/northwood-labs/download-asset) which simplifies building multi-platform images.
* **AWS Organization Security:** Built a [library + CLI tool](https://github.com/northwood-labs/assume-spoke-role) which simplifies the hub-and-spoke pattern for multi-account orgs.
* **AWS Session Manager:** Built a [TUI](https://en.wikipedia.org/wiki/Text-based_user_interface) for [simplifying connections to SSM-enabled EC2 instances](https://github.com/northwood-labs/ssm-shell) using your Terminal.

## Examples of Technical Documentation

Here are examples of my public-facing documentation:

* [Setting up macOS for development](https://github.com/northwood-labs/macos-for-development/wiki)
* [Local AWS Lambda environments (with Go)](https://github.com/northwood-labs/local-lambda-environments-with-go)
* [Local development environment (devsec-tools)](https://github.com/northwood-labs/devsec-tools/blob/main/docs/localdev.md)
* [Configuring DataGrip for Valkey (devsec-tools)](https://github.com/northwood-labs/devsec-tools/blob/main/docs/configuring-datagrip-for-valkey.md)
* Diagrams of Artifactory [infrastructure](https://whimsical.com/artifactory-infrastructure-diagram-5MbWJd8BJfMbRWhfZwtHhQ) and [software](https://whimsical.com/artifactory-software-configuration-and-data-flow-QWZnvbv4SXTX2qkmKGZAqp) configuration.
* Diagram of a [secrets-rotation system](https://whimsical.com/tokengen-infrastructure-diagram-and-data-flow-5ZphqPDP826AaPZHKxCKSr).

## Recommendations

See a [selective list of recommendations](https://github.com/skyzyx/resume/blob/master/selected-recommendations.md) from co-workers and peers.

## Patents and Notable Open-Source

* U.S. patent filing, [“System and Methods for User Authentication across Multiple Domains”][US15042104] (US15042104) (2016)
* U.S. patent filing, [“Hive-based Peer-to-Peer Network”][US8103870B2] (US8103870B2) (2007)
* [SimplePie] — An RSS parser for PHP; founded in 2004; integrated into [WordPress] core since 2009. Millions of global users.
* [CloudFusion][AWS SDK for PHP] — A PHP SDK for AWS; founded in 2005; later became the official [AWS SDK for PHP]. Millions of global users.

## Skills

This list is not exhaustive, but these are software and hard skills I leveraged in the roles above which are most relevant to Cloud Engineering, DevOps, and Site Reliability Engineering roles.

DevOps, DevSecOps, [ARM64], [AWS Well-Architected], [Amazon Web Services][AWS], [Ansible], [Artifactory], [Bash], [CIS], [CentOS], [CloudFormation][AWS CloudFormation], [CloudWatch][Amazon CloudWatch], [Control Tower][AWS Control Tower], [Docker], [ECS][Amazon ECS], [GitHub Actions], [GitHub Enterprise], [Git], [Go], [JavaScript], [Lambda][AWS Lambda], [Nginx], [OpenTelemetry], [OpenTofu], [PHP], [Packer], [Python], [Redis], [SDKs][AWS SDKs], [Terraform], [Traefik], architecture, automation, disaster recovery, error budget, multi-platform, performance, reliability, scalability, scripting, security, troubleshooting, uptime.

## Education

_Silicon Valley College_ (now [Carrington College]), San Jose, CA. Bachelor of Arts, _Design and Visualization_

[Alpine Linux]: https://alpinelinux.org
[Amazon ACM]: https://aws.amazon.com/certificate-manager/
[Amazon CloudFront]: https://aws.amazon.com/cloudfront/
[Amazon CloudWatch]: https://aws.amazon.com/cloudwatch/
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
[Datadog]: https://www.datadoghq.com
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
[Grafana]: https://grafana.com
[GraphQL]: https://graphql.org
[JavaScript]: https://developer.mozilla.org/en-US/docs/Web/JavaScript
[Jenkins]: https://www.jenkins.io
[Jira]: https://www.atlassian.com/software/jira
[JWT]: https://jwt.io
[kubectl]: https://github.com/kubernetes/kubectl
[Kubernetes]: https://kubernetes.io
[LinkedIn]: https://www.linkedin.com/in/rparman/
[McGraw Hill]: https://www.crunchbase.com/organization/mcgraw-hill-education
[New Relic]: https://newrelic.com
[NFS]: https://en.wikipedia.org/wiki/Network_File_System
[Nginx]: https://www.nginx.com
[Northwood Labs]: https://www.crunchbase.com/organization/northwood-labs
[OpenTelemetry]: https://opentelemetry.io
[OpenTofu]: https://opentofu.org
[Packer]: https://packer.io
[PHP-FIG]: http://www.php-fig.org
[PHP]: https://www.php.net
[Python]: https://www.python.org
[Redis]: https://redis.io
[REST]: https://martinfowler.com/articles/richardsonMaturityModel.html
[SB2]: https://www.mheducation.com/news-media/press-releases/mcgraw-hill-connect-unveils-smartbook.html
[SimplePie]: http://simplepie.org
[SLO]: https://sre.google/sre-book/service-level-objectives/
[SRE]: https://sre.google/in-conversation/
[Stack Overflow]: https://stackoverflow.com/users/228514/ryan-parman
[Stripe]: https://stripe.com
[Subversion]: https://subversion.apache.org
[Terraform]: https://www.terraform.io
[Traefik]: https://traefik.io
[twelve-factor applications]: https://12factor.net
[Ubuntu]: https://ubuntu.com
[US15042104]: https://patents.google.com/patent/US20160241536A1/en?inventor=Ryan+Parman
[US8103870B2]: https://patents.google.com/patent/US8103870B2/en?inventor=Ryan+Parman
[Vagrant]: https://www.vagrantup.com
[WarpShare]: https://www.crunchbase.com/organization/warpshare
[WePay]: https://www.crunchbase.com/organization/wepay
[WordPress]: https://wordpress.org
[XSLT]: https://developer.mozilla.org/en-US/docs/Web/XSLT
