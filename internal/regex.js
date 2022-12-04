const removePrefix = (prefix, str) =>
  str.startsWith(prefix) ? str.substr(prefix.length, str.length) : str;
const popWhileMatch = (match, str) => {
  let i = 0;
  while (i < str.length && str[i].match(match) != null) i++;
  return [str.substr(0, i), str.substr(i, str.length)];
};
const nameR = /\w/g;
const whitespace = /\s/g;
const type = /[\w*\.]/g;
const primitives = [
  "float32",
  "float64",
  "string",
  "bool",
  "rune",
  "byte",
  "error",
];
["8", "16", "32", "64"].forEach(size =>
  primitives.push(`int${size}`, `uint${size}`)
);
[...primitives].forEach(primitive => primitives.push(`*${primitive}`));
[...primitives].forEach(primitive => primitives.push(`[]${primitive}`));
console.log(primitives);

const tyPrefix = /[\[\]*]/g;
const packageNonNativeType = (package, ty) => {
  let prefix = "";
  if (ty.includes(".")) return ty;
  [prefix, ty] = popWhileMatch(tyPrefix, ty);
  if (!primitives.includes(ty)) {
    return `${prefix}${package}.${ty}`;
  }
  return prefix + ty;
};

const parseHead = (package, head) => {
  console.log(head);
  head = removePrefix("func (", head);
  [_, head] = popWhileMatch(nameR, head);
  [_, head] = popWhileMatch(whitespace, head);
  [_, head] = popWhileMatch(type, head);
  head = removePrefix(")", head);
  [_, head] = popWhileMatch(whitespace, head);
  let name = "";
  [name, head] = popWhileMatch(nameR, head);
  [_, head] = popWhileMatch(whitespace, head);
  head = removePrefix("(", head);
  const args = [];

  let arg = "",
    ty = "";
  while (!head.startsWith(")")) {
    [arg, head] = popWhileMatch(nameR, head);
    if (head.startsWith(",")) {
      head = removePrefix(",", head);
      [_, head] = popWhileMatch(whitespace, head);
      if (arg == "") continue;
      console.log("Push args", arg, args);
      args.push([arg, ""]);
      continue;
    }
    [_, head] = popWhileMatch(whitespace, head);
    [ty, head] = popWhileMatch(type, head);
    ty = packageNonNativeType(package, ty);
    args.push([arg, ty]);
    for (let i = 0; i < args.length; i++) {
      if (args[i][1] == "") {
        args[i][1] = ty;
      }
    }
  }
  head = removePrefix(")", head);

  [_, head] = popWhileMatch(whitespace, head);
  let returns = [];
  if (head.startsWith("(")) {
    head = removePrefix("(", head);
    while (!head.startsWith(")")) {
      [ty, head] = popWhileMatch(type, head);
      if (!head.startsWith(",")) {
        [_, head] = popWhileMatch(whitespace, head);
        [ty, head] = popWhileMatch(type, head);
      } else {
        head = removePrefix(",", head);
        [_, head] = popWhileMatch(whitespace, head);
      }
      ty = packageNonNativeType(package, ty);
      returns.push(ty);
    }
    head = removePrefix(")", head);
  } else {
    [ty, head] = popWhileMatch(type, head);
    ty = packageNonNativeType(package, ty);
    returns.push(ty);
  }
  return {
    name,
    noreturn: returns.length < 2,
    noerr: returns[returns.length - 1] != "error",
    returns: returns[0],
    args,
    onErrFactory: `empty[${returns[0]}]`,
  };
};

const toFlow = (name, package, type, heads) => {
  return {
    base: name,
    baseName: name.charAt(0).toLowerCase() + name.slice(1),
    wraps: type,
    wrapName: name.charAt(0).toLowerCase() + name.slice(1),
    functions: heads
      .split("\n")
      .filter(line => line.trim() != "")
      .map(head => parseHead(package, head)),
  };
};

console.log(
  JSON.stringify(
    toFlow(
      "HttpClient",
      "http",
      "*http.Client",
      `func (c *Client) CloseIdleConnections()
func (c *Client) Post(url, contentType string, body io.Reader) (resp *Response, err error)
func (c *Client) Do(req *Request) (*Response, error)
func (c *Client) Get(url string) (resp *Response, err error)
func (c *Client) Head(url string) (resp *Response, err error)
func (c *Client) PostForm(url string, data url.Values) (resp *Response, err error)`
    )
  )
);
