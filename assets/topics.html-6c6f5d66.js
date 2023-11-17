import{_ as p,r as o,o as i,c,a as s,b as a,w as r,d as n,e}from"./app-0a598c04.js";const u={},d=e(`<h1 id="export-topics" tabindex="-1"><a class="header-anchor" href="#export-topics" aria-hidden="true">#</a> Export Topics</h1><p>Export Topics metadata to different formats.</p><p>Includes Topic configuration and Topic ACLs.</p><div class="language-bash" data-ext="sh"><pre class="language-bash"><code>cctools <span class="token builtin class-name">export</span> topics --help\`
</code></pre></div><div class="language-bash" data-ext="sh"><pre class="language-bash"><code> Command to <span class="token builtin class-name">export</span> Topics information.

Usage:
  cctools <span class="token builtin class-name">export</span> topics <span class="token punctuation">[</span>flags<span class="token punctuation">]</span> 

Flags:
  -h, <span class="token parameter variable">--help</span>   <span class="token builtin class-name">help</span> <span class="token keyword">for</span> topics

Global Flags:
  -c, <span class="token parameter variable">--config</span> string   config <span class="token function">file</span> 
  -o, <span class="token parameter variable">--output</span> string   Output format. Possible values: json, yaml, hcl, cfk, clink
</code></pre></div><h2 id="configuration" tabindex="-1"><a class="header-anchor" href="#configuration" aria-hidden="true">#</a> Configuration</h2><ul><li>Using Topic Exporter Configuration to exclude some topics.</li></ul><p>All topics names containing <code>_confluent</code> will be excluded.</p><div class="language-yaml line-numbers-mode" data-ext="yml"><pre class="language-yaml"><code><span class="token key atrule">export</span><span class="token punctuation">:</span> 
  <span class="token key atrule">topics</span><span class="token punctuation">:</span>
    <span class="token key atrule">exclude</span><span class="token punctuation">:</span> _confluent 
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><ul><li>Topic ACLs - Principals Mapping</li></ul><p>All the Topic ACLs where <code>principal: User:test</code> will be created as <code>principal: User:sa-xyroox</code> on the Destination.</p><div class="language-yaml line-numbers-mode" data-ext="yml"><pre class="language-yaml"><code><span class="token key atrule">principals</span><span class="token punctuation">:</span>
  <span class="token punctuation">-</span> <span class="token key atrule">&quot;test&quot;</span><span class="token punctuation">:</span> <span class="token string">&quot;sa-xyroox&quot;</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div>`,12),k=e(`<h2 id="export-format" tabindex="-1"><a class="header-anchor" href="#export-format" aria-hidden="true">#</a> Export format</h2><p>Required <code>--output</code></p><p>Output format:</p><ul><li>JSON:</li></ul><div class="language-bash line-numbers-mode" data-ext="sh"><pre class="language-bash"><code>  cctools <span class="token builtin class-name">export</span> topics <span class="token parameter variable">--output</span> json <span class="token parameter variable">--config</span> config.yaml
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><ul><li>YAML:</li></ul><div class="language-bash" data-ext="sh"><pre class="language-bash"><code>  cctools <span class="token builtin class-name">export</span> topics <span class="token parameter variable">--output</span> yaml <span class="token parameter variable">--config</span> config.yaml
</code></pre></div><ul><li>CFK(YML):</li></ul><div class="language-bash" data-ext="sh"><pre class="language-bash"><code>  cctools <span class="token builtin class-name">export</span> topics <span class="token parameter variable">--output</span> cfk <span class="token parameter variable">--config</span> config.yaml
</code></pre></div><ul><li>CLINK(SH):</li></ul><div class="language-bash" data-ext="sh"><pre class="language-bash"><code>cctools <span class="token builtin class-name">export</span> topics <span class="token parameter variable">--output</span> clink <span class="token parameter variable">--config</span> config.yaml
</code></pre></div><ul><li>HCL(TFVARS):</li></ul><div class="language-bash" data-ext="sh"><pre class="language-bash"><code>cctools <span class="token builtin class-name">export</span> topics <span class="token parameter variable">--output</span> hcl <span class="token parameter variable">--config</span> config.yaml
</code></pre></div><h2 id="example" tabindex="-1"><a class="header-anchor" href="#example" aria-hidden="true">#</a> Example</h2><div class="language-bash" data-ext="sh"><pre class="language-bash"><code>cctools <span class="token builtin class-name">export</span> topics <span class="token parameter variable">--output</span> json <span class="token parameter variable">--config</span> config.yaml
</code></pre></div><ul><li>Source cluster configuration (config.yaml)</li><li>Exclude consiguration.</li></ul><p>Exporter will create a <code>JSON</code> file per topic selected.</p><p>Topic selected : <code>demo.topic</code>. <code>demo.topic.json</code> file under <code>output/topics/json</code> folder.</p><div class="language-json line-numbers-mode" data-ext="json"><pre class="language-json"><code><span class="token punctuation">{</span>
 <span class="token property">&quot;name&quot;</span><span class="token operator">:</span> <span class="token string">&quot;demo.topic&quot;</span><span class="token punctuation">,</span>
 <span class="token property">&quot;partitions&quot;</span><span class="token operator">:</span> <span class="token number">4</span><span class="token punctuation">,</span>
 <span class="token property">&quot;replicationFactor&quot;</span><span class="token operator">:</span> <span class="token number">3</span><span class="token punctuation">,</span>
 <span class="token property">&quot;minIsr&quot;</span><span class="token operator">:</span> <span class="token string">&quot;2&quot;</span><span class="token punctuation">,</span>
 <span class="token property">&quot;retentionTime&quot;</span><span class="token operator">:</span> <span class="token string">&quot;604800000&quot;</span><span class="token punctuation">,</span>
 <span class="token property">&quot;configs&quot;</span><span class="token operator">:</span> <span class="token punctuation">[</span>
  <span class="token punctuation">{</span>
   <span class="token property">&quot;name&quot;</span><span class="token operator">:</span> <span class="token string">&quot;index.interval.bytes&quot;</span><span class="token punctuation">,</span>
   <span class="token property">&quot;value&quot;</span><span class="token operator">:</span> <span class="token string">&quot;4096&quot;</span>
  <span class="token punctuation">}</span><span class="token punctuation">,</span>
  ...
 <span class="token punctuation">]</span><span class="token punctuation">,</span>
 
 <span class="token property">&quot;acls&quot;</span><span class="token operator">:</span> <span class="token punctuation">[</span>
  <span class="token punctuation">{</span>
   <span class="token property">&quot;principal&quot;</span><span class="token operator">:</span> <span class="token string">&quot;User:test&quot;</span><span class="token punctuation">,</span>
   <span class="token property">&quot;host&quot;</span><span class="token operator">:</span> <span class="token string">&quot;*&quot;</span><span class="token punctuation">,</span>
   <span class="token property">&quot;operation&quot;</span><span class="token operator">:</span> <span class="token string">&quot;ALL&quot;</span><span class="token punctuation">,</span>
   <span class="token property">&quot;permission&quot;</span><span class="token operator">:</span> <span class="token string">&quot;ALLOW&quot;</span><span class="token punctuation">,</span>
   <span class="token property">&quot;resourceType&quot;</span><span class="token operator">:</span> <span class="token string">&quot;TOPIC&quot;</span><span class="token punctuation">,</span>
   <span class="token property">&quot;resourceName&quot;</span><span class="token operator">:</span> <span class="token string">&quot;demo.topic&quot;</span><span class="token punctuation">,</span>
   <span class="token property">&quot;patternType&quot;</span><span class="token operator">:</span> <span class="token string">&quot;LITERAL&quot;</span>
  <span class="token punctuation">}</span>
 <span class="token punctuation">]</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="yaml" tabindex="-1"><a class="header-anchor" href="#yaml" aria-hidden="true">#</a> YAML</h3><p>OutPut: <code>&lt;output_path&gt; / &lt;cluster_ID&gt;_&lt;resource&gt;.yaml</code></p><p>Output Sample for <code>topics</code> resource:</p><div class="language-yaml line-numbers-mode" data-ext="yml"><pre class="language-yaml"><code><span class="token punctuation">-</span> <span class="token key atrule">name</span><span class="token punctuation">:</span> demo.topic
  <span class="token key atrule">partitions</span><span class="token punctuation">:</span> <span class="token number">1</span>
  <span class="token key atrule">replicationfactor</span><span class="token punctuation">:</span> <span class="token number">3</span>
  <span class="token key atrule">configs</span><span class="token punctuation">:</span>
  <span class="token punctuation">-</span> <span class="token key atrule">name</span><span class="token punctuation">:</span> cleanup.policy
    <span class="token key atrule">value</span><span class="token punctuation">:</span> compact
  <span class="token punctuation">-</span> <span class="token key atrule">name</span><span class="token punctuation">:</span> compression.type
    <span class="token key atrule">value</span><span class="token punctuation">:</span> producer
  <span class="token punctuation">...</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="cfk" tabindex="-1"><a class="header-anchor" href="#cfk" aria-hidden="true">#</a> CFK</h3>`,24),m={href:"https://docs.confluent.io/operator/current/co-manage-topics.html#create-ak-topic",target:"_blank",rel:"noopener noreferrer"},v=s("code",null,"KafkaTopic",-1),b=e(`<p>Output: Exporter will create a <code>&lt;output_path&gt;/&lt;clusterid&gt;_&lt;resource&gt;_&lt;topicName&gt;.yml</code> file for each Topic.</p><p>Example:</p><div class="language-yaml line-numbers-mode" data-ext="yml"><pre class="language-yaml"><code><span class="token key atrule">apiVersion</span><span class="token punctuation">:</span> platform.confluent.io/v1beta1
<span class="token key atrule">kind</span><span class="token punctuation">:</span> KafkaTopic
<span class="token key atrule">metadata</span><span class="token punctuation">:</span>
  <span class="token key atrule">name</span><span class="token punctuation">:</span> demo.topic
  <span class="token key atrule">namespace</span><span class="token punctuation">:</span> confluent
<span class="token key atrule">spec</span><span class="token punctuation">:</span>
  <span class="token key atrule">replicas</span><span class="token punctuation">:</span> <span class="token number">3</span>
  <span class="token key atrule">partitionCount</span><span class="token punctuation">:</span> <span class="token number">6</span>
  <span class="token key atrule">configs</span><span class="token punctuation">:</span>
    <span class="token key atrule">cleanup.policy</span><span class="token punctuation">:</span> delete
    <span class="token key atrule">compression.type</span><span class="token punctuation">:</span> producer
    <span class="token punctuation">...</span>
  <span class="token key atrule">kafkaRestClassRef</span><span class="token punctuation">:</span>
    <span class="token key atrule">name</span><span class="token punctuation">:</span> kafka

</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="cluster-link" tabindex="-1"><a class="header-anchor" href="#cluster-link" aria-hidden="true">#</a> Cluster Link</h3>`,4),g=s("code",null,"topics",-1),h={href:"https://docs.confluent.io/cloud/current/multi-cloud/overview.html",target:"_blank",rel:"noopener noreferrer"},f=e(`<p>Output: The export will generate:</p><ul><li>Cluster Link creation script (.sh), including topic mirrors from selected <code>topics</code> if <code>autocreate</code> is <code>false</code></li><li>Topic promotion script (.sh)</li><li>Clean up script (.sh)</li><li>Cluster Link configuration file (.properties), including <code>auto.create.mirror.topics.filters</code> from selected <code>topics</code></li></ul><h3 id="hcl" tabindex="-1"><a class="header-anchor" href="#hcl" aria-hidden="true">#</a> HCL</h3><p>Exporting Topics to HCL will generate: <code>output/topics/tfvars/topics.tfvars</code>.</p><div class="language-json line-numbers-mode" data-ext="json"><pre class="language-json"><code>environment = <span class="token string">&quot;&lt;ENV_ID&gt;&quot;</span>

cluster = <span class="token string">&quot;&lt;CLUSTER_ID&gt;&quot;</span>

rbac_enabled = <span class="token boolean">false</span>

serv_account = <span class="token punctuation">{</span>
  name = <span class="token string">&quot;&lt;SERVICE_ACCOUNT&gt;&quot;</span>
  role = <span class="token string">&quot;CloudClusterAdmin&quot;</span>
<span class="token punctuation">}</span>
topics = <span class="token punctuation">[</span><span class="token punctuation">{</span>
  name       = <span class="token string">&quot;orders&quot;</span>
  partitions = <span class="token number">12</span>
  config = <span class="token punctuation">{</span>
    <span class="token string">&quot;cleanup.policy&quot;</span>                          = <span class="token string">&quot;delete&quot;</span>
    <span class="token string">&quot;compression.type&quot;</span>                        = <span class="token string">&quot;producer&quot;</span>
    ...
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div>`,5),q={href:"https://github.com/mcolomerc/terraform-confluent-topics",target:"_blank",rel:"noopener noreferrer"},y=s("h3",{id:"excel",tabindex:"-1"},[s("a",{class:"header-anchor",href:"#excel","aria-hidden":"true"},"#"),n(" Excel")],-1),x=s("p",null,[n("Output: "),s("code",null,"<output_path>/<cluster_ID>_<resource>.xlsx")],-1);function _(C,T){const l=o("RouterLink"),t=o("ExternalLinkIcon");return i(),c("div",null,[d,s("p",null,[a(l,{to:"/commands/config/"},{default:r(()=>[n("Configuration")]),_:1})]),k,s("p",null,[n("From the selected resources export to "),s("a",m,[n("Confluent For Kubernetes (CFK) Custom Resources"),a(t)]),n(),v,n(".")]),b,s("p",null,[n("From the selected "),g,n(" export "),s("a",h,[n("Confluent Cloud Cluster Link"),a(t)]),n(" scripts and configuration.")]),f,s("p",null,[n("The output could be used with "),s("a",q,[n("Terraform Topics Module"),a(t)]),n(" to create topics on Confluent cloud destination cluster.")]),y,x])}const E=p(u,[["render",_],["__file","topics.html.vue"]]);export{E as default};
