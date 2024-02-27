#!/usr/bin/env node
import "source-map-support/register";
import { App } from "aws-cdk-lib";
import { Code } from "aws-cdk-lib/aws-lambda";
import {
  TexitStatefulResourcesStack,
  TexitApiStack,
  TexitDiscordBotStack,
  TexitWorkflowsStack,
} from "texit-constructs";

const app = new App();

const resources = new TexitStatefulResourcesStack(
  app,
  "TexitStatefulResourcesStack",
  {
    configsPath: "assets/config/test",
  }
);

const workflows = new TexitWorkflowsStack(app, "TexitWorkflowsStack", {
  binary: Code.fromAsset("assets/bin/texit-activities"),
  configBucket: resources.configBucket,
  nodeTable: resources.nodesTable,
  executionTable: resources.executionsTable,
  notifierTopic: resources.notifierTopic,
});

const texit = new TexitApiStack(app, "TexitApiStack", {
  binary: Code.fromAsset("assets/bin/texit"),
  configBucket: resources.configBucket,
  nodeTable: resources.nodesTable,
  executionTable: resources.executionsTable,
  provisionNodeWorkflow: workflows.provisionNodeWorkflow,
  deployNodeWorkflow: workflows.deprovisionNodeWorkflow,
  notifierTopic: resources.notifierTopic,
});

new TexitDiscordBotStack(app, "TexitDiscordBotStack", {
  botBinary: Code.fromAsset("assets/bin/texit-discord"),
  callbackBinary: Code.fromAsset("assets/bin/texit-discord-callback"),
  configBucket: resources.configBucket,
  configObject: "bot-config.yaml",
  texitEndpoint: texit.api.url!,
  callbackTopic: resources.notifierTopic,
});
