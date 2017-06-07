package main

const configuration = `<flow-definition plugin="workflow-job@2.10">
  <actions/>
  <description>{{ .Description }}</description>
  {{- with .Parameters }}
  <keepDependencies>{{ .KeepDependencies }}</keepDependencies>
  <properties>
    <com.dabsquared.gitlabjenkins.connection.GitLabConnectionProperty plugin="gitlab-plugin@1.4.5">
      <gitLabConnection>{{ .GitLabConnection }}</gitLabConnection>
    </com.dabsquared.gitlabjenkins.connection.GitLabConnectionProperty>
    <org.jenkinsci.plugins.workflow.job.properties.PipelineTriggersJobProperty>
      <triggers>
        <com.dabsquared.gitlabjenkins.GitLabPushTrigger plugin="gitlab-plugin@1.4.5">
          <spec></spec>
          <triggerOnPush>{{ .TriggerOnPush }}</triggerOnPush>
          <triggerOnMergeRequest>{{ .TriggerOnMergeRequest }}</triggerOnMergeRequest>
          <triggerOpenMergeRequestOnPush>{{ .TriggerOpenMergeRequestOnPush }}</triggerOpenMergeRequestOnPush>
          <triggerOnNoteRequest>{{ .TriggerOnNoteRequest }}</triggerOnNoteRequest>
          <noteRegex>{{ .NoteRegex }}</noteRegex>
          <ciSkip>{{ .CISkip }}</ciSkip>
          <skipWorkInProgressMergeRequest>{{ .SkipWorkInProgressMergeRequest }}</skipWorkInProgressMergeRequest>
          <setBuildDescription>{{ .SetBuildDescription }}</setBuildDescription>
          <branchFilterType>{{ .BranchFilterType }}</branchFilterType>
          <includeBranchesSpec></includeBranchesSpec>
          <excludeBranchesSpec></excludeBranchesSpec>
          <targetBranchRegex></targetBranchRegex>
          <secretToken>{{ .SecretToken }}</secretToken>
        </com.dabsquared.gitlabjenkins.GitLabPushTrigger>
      </triggers>
    </org.jenkinsci.plugins.workflow.job.properties.PipelineTriggersJobProperty>
  </properties>
  <definition class="org.jenkinsci.plugins.workflow.cps.CpsScmFlowDefinition" plugin="workflow-cps@2.30">
    <scm class="hudson.plugins.git.GitSCM" plugin="git@3.3.0">
      <configVersion>2</configVersion>
      <userRemoteConfigs>
        <hudson.plugins.git.UserRemoteConfig>
          <url>https://git.utenze.bankit.it/gestconf/gcportal-web.git</url>
        </hudson.plugins.git.UserRemoteConfig>
      </userRemoteConfigs>
      <branches>
        <hudson.plugins.git.BranchSpec>
          <name>*/master</name>
        </hudson.plugins.git.BranchSpec>
      </branches>
      <doGenerateSubmoduleConfigurations>false</doGenerateSubmoduleConfigurations>
      <submoduleCfg class="list"/>
      <extensions/>
    </scm>
    <scriptPath>Jenkinsfile</scriptPath>
    <lightweight>true</lightweight>
  </definition>
  <triggers/>
  {{- end }}
</flow-definition>
`
