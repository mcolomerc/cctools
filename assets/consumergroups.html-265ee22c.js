import{_ as s,o as a,c as n,e}from"./app-0a598c04.js";const o={},r=e(`<h1 id="consumer-groups" tabindex="-1"><a class="header-anchor" href="#consumer-groups" aria-hidden="true">#</a> Consumer Groups</h1><p>Consumer groups are a way of grouping Kafka consumers together to consume a topic. Each consumer in a group will consume from a unique subset of partitions in the topic. This allows you to horizontally scale your consumers while still maintaining the ordering guarantees of a single partition.</p><p>Exporting consumer groups will generate a file per consumer group.</p><div class="language-bash" data-ext="sh"><pre class="language-bash"><code>cctools <span class="token builtin class-name">export</span> consumer-groups <span class="token parameter variable">--help</span>
</code></pre></div><p>Usage:</p><div class="language-bash" data-ext="sh"><pre class="language-bash"><code> Command to <span class="token builtin class-name">export</span> Consumer Group information.

Usage:
  cctools <span class="token builtin class-name">export</span> consumer-groups <span class="token punctuation">[</span>flags<span class="token punctuation">]</span> 

Flags:
  -h, <span class="token parameter variable">--help</span>   <span class="token builtin class-name">help</span> <span class="token keyword">for</span> consumer-groups

Global Flags:
  -c, <span class="token parameter variable">--config</span> string   config <span class="token function">file</span>
  -o, <span class="token parameter variable">--output</span> string   Output format. Possible values: json, yaml, hcl, cfk, clink
</code></pre></div>`,6),t=[r];function c(l,p){return a(),n("div",null,t)}const u=s(o,[["render",c],["__file","consumergroups.html.vue"]]);export{u as default};
