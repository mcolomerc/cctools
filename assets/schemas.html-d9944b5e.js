import{_ as a,o as s,c as n,e}from"./app-0a598c04.js";const t={},o=e(`<h1 id="export-schemas" tabindex="-1"><a class="header-anchor" href="#export-schemas" aria-hidden="true">#</a> Export Schemas</h1><div class="language-bash" data-ext="sh"><pre class="language-bash"><code>cctools <span class="token builtin class-name">export</span> schemas <span class="token parameter variable">--help</span>
</code></pre></div><div class="language-bash" data-ext="sh"><pre class="language-bash"><code> Command to <span class="token builtin class-name">export</span> Schemas information.

Usage:
  cctcctools <span class="token builtin class-name">export</span> schemas <span class="token punctuation">[</span>flags<span class="token punctuation">]</span> 

Flags:
  -h, <span class="token parameter variable">--help</span>   <span class="token builtin class-name">help</span> <span class="token keyword">for</span> schemas

Global Flags:
  -c, <span class="token parameter variable">--config</span> string   config <span class="token function">file</span>  
  -o, <span class="token parameter variable">--output</span> string   Output format. Possible values: json, yaml, hcl, cfk, clink
</code></pre></div><p>Output format:</p><ul><li>JSON: <code>cctools export schemas --output json --config config.yaml</code></li><li>YAML: <code>cctools export schemas --output yaml --config config.yaml</code></li><li>CFK(YML): <code>cctools export schemas --output cfk --config config.yaml</code></li><li>Excel(XLS): <code>cctools export schemas --output excel --config config.yaml</code></li></ul><h2 id="configuration" tabindex="-1"><a class="header-anchor" href="#configuration" aria-hidden="true">#</a> Configuration</h2><p>Configure Subject export: <code>all</code> subject versions or only the <code>latest</code> version.</p><div class="language-yaml line-numbers-mode" data-ext="yml"><pre class="language-yaml"><code><span class="token key atrule">export</span><span class="token punctuation">:</span> 
  <span class="token key atrule">schemas</span><span class="token punctuation">:</span> 
    <span class="token key atrule">version</span><span class="token punctuation">:</span> latest  <span class="token comment"># default: all </span>
    <span class="token key atrule">subjects</span><span class="token punctuation">:</span>
      <span class="token key atrule">version</span><span class="token punctuation">:</span> latest <span class="token comment"># default: all </span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="export-format" tabindex="-1"><a class="header-anchor" href="#export-format" aria-hidden="true">#</a> Export format</h2><h3 id="json" tabindex="-1"><a class="header-anchor" href="#json" aria-hidden="true">#</a> JSON</h3><p>Schemas:</p><p>Output path: <code>&lt;export.output.path&gt;/schemas/json</code></p><div class="language-json line-numbers-mode" data-ext="json"><pre class="language-json"><code><span class="token punctuation">{</span>
 <span class="token property">&quot;subject&quot;</span><span class="token operator">:</span> <span class="token string">&quot;customer-value&quot;</span><span class="token punctuation">,</span>
 <span class="token property">&quot;version&quot;</span><span class="token operator">:</span> <span class="token number">1</span><span class="token punctuation">,</span>
 <span class="token property">&quot;id&quot;</span><span class="token operator">:</span> <span class="token number">100011</span><span class="token punctuation">,</span>
 <span class="token property">&quot;schema&quot;</span><span class="token operator">:</span> <span class="token string">&quot;{\\&quot;type\\&quot;:\\&quot;record\\&quot;,\\&quot;name\\&quot;:\\&quot;Customer\\&quot;,\\&quot;fields\\&quot;:...&quot;</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Subjects:</p><p><code>&lt;export.output.path&gt;/subjects/json</code></p><div class="language-json line-numbers-mode" data-ext="json"><pre class="language-json"><code><span class="token punctuation">{</span>
 <span class="token property">&quot;subject&quot;</span><span class="token operator">:</span> <span class="token string">&quot;customer-value&quot;</span><span class="token punctuation">,</span>
 <span class="token property">&quot;version&quot;</span><span class="token operator">:</span> <span class="token number">1</span><span class="token punctuation">,</span>
 <span class="token property">&quot;id&quot;</span><span class="token operator">:</span> <span class="token number">100011</span><span class="token punctuation">,</span>
 <span class="token property">&quot;schema&quot;</span><span class="token operator">:</span> <span class="token string">&quot;{\\&quot;type\\&quot;:\\&quot;record\\&quot;,\\&quot;name\\&quot;:\\&quot;Customer\\&quot;,\\&quot;fields\\&quot;:[...&quot;</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="yaml" tabindex="-1"><a class="header-anchor" href="#yaml" aria-hidden="true">#</a> YAML</h3><p>Schemas:</p><p><code>&lt;export.output.path&gt;/schemas/yaml</code></p><div class="language-yaml line-numbers-mode" data-ext="yml"><pre class="language-yaml"><code><span class="token key atrule">subject</span><span class="token punctuation">:</span> payment<span class="token punctuation">-</span>value
<span class="token key atrule">version</span><span class="token punctuation">:</span> <span class="token number">1</span>
<span class="token key atrule">id</span><span class="token punctuation">:</span> <span class="token number">100064</span>
<span class="token key atrule">schemaType</span><span class="token punctuation">:</span> <span class="token string">&quot;&quot;</span>
<span class="token key atrule">schema</span><span class="token punctuation">:</span> <span class="token string">&#39;{&quot;type&quot;:&quot;record&quot;,&quot;name&quot;:&quot;Payment&quot;,&quot;namespace&quot;:&quot;io.confluent.examples.clients.basicavro&quot;,&quot;fields&quot;:[{...&#39;</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Subjects:</p><p><code>&lt;export.output.path&gt;/subjects/yaml</code></p><div class="language-yaml line-numbers-mode" data-ext="yml"><pre class="language-yaml"><code><span class="token key atrule">subject</span><span class="token punctuation">:</span> payment<span class="token punctuation">-</span>value
<span class="token key atrule">version</span><span class="token punctuation">:</span> <span class="token number">1</span>
<span class="token key atrule">id</span><span class="token punctuation">:</span> <span class="token number">100064</span>
<span class="token key atrule">schemaType</span><span class="token punctuation">:</span> <span class="token string">&quot;&quot;</span>
<span class="token key atrule">schema</span><span class="token punctuation">:</span> <span class="token string">&#39;{&quot;type&quot;:&quot;record&quot;,&quot;name&quot;:&quot;Payment&quot;,&quot;namespace&quot;:&quot;io.confluent.examples.clients.basicavro&quot;,&quot;fields&quot;:[{&quot;name&quot;:&quot;id&quot;,&quot;type&quot;:&quot;string&quot;},{&quot;name&quot;:&quot;amount&quot;,&quot;type&quot;:&quot;double&quot;},{&quot;name&quot;:&quot;email&quot;,&quot;type&quot;:&quot;string&quot;}]}&#39;</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="cfk" tabindex="-1"><a class="header-anchor" href="#cfk" aria-hidden="true">#</a> CFK</h3><p>CFK export configuration:</p><div class="language-yaml line-numbers-mode" data-ext="yml"><pre class="language-yaml"><code><span class="token key atrule">export</span><span class="token punctuation">:</span> 
  <span class="token key atrule">schemas</span><span class="token punctuation">:</span> 
    <span class="token key atrule">version</span><span class="token punctuation">:</span> latest  <span class="token comment"># default: all </span>
    <span class="token key atrule">subjects</span><span class="token punctuation">:</span>
      <span class="token key atrule">version</span><span class="token punctuation">:</span> latest <span class="token comment"># default: all </span>
  <span class="token key atrule">cfk</span><span class="token punctuation">:</span>
    <span class="token key atrule">namespace</span><span class="token punctuation">:</span> confluent  
    <span class="token key atrule">kafkarestclass</span><span class="token punctuation">:</span> kafka 
    <span class="token key atrule">schemaRegistryClusterRef</span><span class="token punctuation">:</span> schemaregistry 
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Output: <code>&lt;export.output.path&gt;/schemas/cfk</code></p><p>Schema:</p><div class="language-yaml line-numbers-mode" data-ext="yml"><pre class="language-yaml"><code><span class="token key atrule">apiVersion</span><span class="token punctuation">:</span> platform.confluent.io/v1beta1
<span class="token key atrule">kind</span><span class="token punctuation">:</span> Schema
<span class="token key atrule">metadata</span><span class="token punctuation">:</span>
  <span class="token key atrule">name</span><span class="token punctuation">:</span> schema_name<span class="token punctuation">-</span>value
  <span class="token key atrule">namespace</span><span class="token punctuation">:</span> confluent
<span class="token key atrule">spec</span><span class="token punctuation">:</span>
  <span class="token key atrule">data</span><span class="token punctuation">:</span>
    <span class="token key atrule">configRef</span><span class="token punctuation">:</span> schema_name<span class="token punctuation">-</span>value<span class="token punctuation">-</span>config
    <span class="token key atrule">format</span><span class="token punctuation">:</span> avro
  <span class="token key atrule">schemaRegistryClusterRef</span><span class="token punctuation">:</span>
    <span class="token key atrule">name</span><span class="token punctuation">:</span> schemaregistry
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Config Map:</p><div class="language-yaml line-numbers-mode" data-ext="yml"><pre class="language-yaml"><code><span class="token key atrule">apiVersion</span><span class="token punctuation">:</span> v1
<span class="token key atrule">kind</span><span class="token punctuation">:</span> ConfigMap
<span class="token key atrule">metadata</span><span class="token punctuation">:</span>
  <span class="token key atrule">name</span><span class="token punctuation">:</span> schema_name<span class="token punctuation">-</span>value<span class="token punctuation">-</span>config
  <span class="token key atrule">namespace</span><span class="token punctuation">:</span> confluent
<span class="token key atrule">data</span><span class="token punctuation">:</span>
  <span class="token key atrule">schema</span><span class="token punctuation">:</span> <span class="token string">&#39;{&quot;type&quot;:&quot;record&quot;,&quot;name&quot;:&quot;record&quot;,&quot;namespace&quot;:&quot;org.apache.flink.avro.generated&quot;,&quot;fields&quot;:...&#39;</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div>`,31),p=[o];function l(c,u){return s(),n("div",null,p)}const r=a(t,[["render",l],["__file","schemas.html.vue"]]);export{r as default};
