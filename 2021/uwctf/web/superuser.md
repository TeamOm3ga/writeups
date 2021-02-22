# web / superuser

> Points: 150

## Question

> A password is all you need.
> [Link](https://docs.google.com/forms/d/e/1FAIpQLSe__FZojKvtjm0P7j3nG51bIatJmC-Q9itpMPpgkLTvHoA9mw/viewform)

## Solution

As I know from possibly or possibly not using this trick in high school,
Google Forms do client-side validation.

We inspect the page and search for a parent element with validation data in an HTML attr:
A few layers of `<div>`s up, we find a `div.m2` with a `data-params` attribute.

```html
<div jsmodel="CP1oW" data-params="%.@.[1093023383,&quot;Password:&quot;,null,0,[[1554011263,[],true,[],[[4,301,[&quot;\\x75\\x77\\x63\\x74\\x66\\x7b\\x63\\x6c\\x69\\x65\\x6e\\x74\\x73\\x69\\x64\\x65\\x76\\x61\\x6c\\x69\\x64\\x61\\x74\\x69\\x6f\\x6e\\x69\\x73\\x62\\x61\\x64\\x5f\\x36\\x30\\x34\\x34\\x65\\x35\\x32\\x62\\x34\\x32\\x38\\x39\\x38\\x33\\x63\\x35\\x7d&quot;],&quot;su: Authentication failure&quot;]],null,null,null,null,null,[null,[]]]],null,null,null,[]],&quot;i1&quot;,&quot;i2&quot;,&quot;i3&quot;,false]" class="m2">
```

Within this, there is a hex-encoded string.
This is the flag, which we can decode by interpreting the `data-params` attribute as an array:

```js
data = JSON.parse('['+$('div.m2').attributes["data-params"].value.substring(4));
flag = data[0][4][0][4][0][2][0]
eval("'"+flag+"'");
// uwctf{clientsidevalidationisbad_6044e52b428983c5}
```

### Flag

`uwctf{clientsidevalidationisbad_6044e52b428983c5}`
