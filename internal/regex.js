const removePrefix = (prefix, str) =>
  str.startsWith(prefix) ? str.substr(prefix.length, str.length) : str;
const popWhileMatch = (match, str) => {
  let i = 0;
  while (i < str.length && str[i].match(match) != null) i++;
  return [str.substr(0, i), str.substr(i, str.length)];
};
const nameR = /\w/g;
const dots = /\./g;
const whitespace = /\s/g;
const type = /[\w*\{\}\[\]\.]/g;
const primitives = [
  "interface{}",
  "any",
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

const parseHead = (flowWraps, package, head) => {
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
    if (head.startsWith("func")) return null;
    [ty, head] = popWhileMatch(type, head);
    ty = packageNonNativeType(package, ty);
    [dotdotdot, ty] = popWhileMatch(dots, ty);
    args.push([arg + dotdotdot, ty]);
    for (let i = 0; i < args.length; i++) {
      if (args[i][1] == "") {
        args[i][1] = ty;
        args[i][0] = args[i][0] + dotdotdot;
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
      if (!head.startsWith(",") && !head.startsWith(")")) {
        [_, head] = popWhileMatch(whitespace, head);
        [ty, head] = popWhileMatch(type, head);
      } else {
        head = removePrefix(",", head);
        [_, head] = popWhileMatch(whitespace, head);
      }
      if (ty != "") {
        ty = packageNonNativeType(package, ty);
        returns.push(ty);
      }
    }
    head = removePrefix(")", head);
  } else {
    [ty, head] = popWhileMatch(type, head);
    if (ty != "") {
      ty = packageNonNativeType(package, ty);
      returns.push(ty);
    }
  }

  if (returns.length > 2) {
    console.log(
      "Function with more than 2 return values not supported",
      returns
    );
    return null;
  }
  if (name.startsWith("Must")) return null;
  console.log(name, returns);
  let extra = {};
  if (returns.length > 0 && returns[0] in flowWraps) {
    let wrap = flowWraps[returns[0]];
    returns[0] = wrap.ty;
    if (wrap.err) extra.onErrFactory = wrap.err;
    if (wrap.success) extra.successFactory = wrap.success;
  }

  return {
    name,
    inflow: true,
    noreturn:
      returns.length == 0 || (returns[0] == "error" && returns.length == 1),
    noerr: returns[returns.length - 1] != "error",
    returns: returns[0],
    args,
    onErrFactory: `empty[${returns[0]}]`,
    ...extra,
  };
};

const toFlow = (name, package, type, flowWraps, heads) => {
  return {
    base: name,
    baseName: name.charAt(0).toLowerCase() + name.slice(1),
    wraps: type,
    wrapName: name.charAt(0).toLowerCase() + name.slice(1),
    functions: heads
      .split("\n")
      .filter(line => line.trim() != "")
      .map(head => parseHead(flowWraps, package, head))
      .filter(line => line),
  };
};

const wraps = {
  "*sqlx.Tx": {
    ty: "*Transactionx",
    success: "TransactionxOf",
    err: "EmptyTransactionxOf",
  },
  "*sql.Stmt": {
    ty: "*Stmt",
    success: "StmtOf",
    err: "EmptyStmtOf",
  },
  "*sqlx.Stmt": {
    ty: "*Stmtx",
    success: "StmtxOf",
    err: "EmptyStmtxOf",
  },
  "sql.Result": {
    ty: "*Result",
    success: "ResultOf",
    err: "EmptyResultOf",
  },
  "sqlx.Result": {
    ty: "*Result",
    success: "ResultOf",
    err: "EmptyResultOf",
  },
};

const tx = () =>
  toFlow(
    "Transaction",
    "sql",
    "*sql.Tx",
    wraps,
    `func (tx *Tx) Commit() error
func (tx *Tx) Exec(query string, args ...any) (Result, error)
func (tx *Tx) ExecContext(ctx context.Context, query string, args ...any) (Result, error)
func (tx *Tx) Prepare(query string) (*Stmt, error)
func (tx *Tx) PrepareContext(ctx context.Context, query string) (*Stmt, error)
func (tx *Tx) Query(query string, args ...any) (*Rows, error)
func (tx *Tx) QueryContext(ctx context.Context, query string, args ...any) (*Rows, error)
func (tx *Tx) QueryRow(query string, args ...any) *Row
func (tx *Tx) QueryRowContext(ctx context.Context, query string, args ...any) *Row
func (tx *Tx) Rollback() error
func (tx *Tx) Stmt(stmt *Stmt) *Stmt
func (tx *Tx) StmtContext(ctx context.Context, stmt *Stmt) *Stmt`
  );

const txx = () => {
  let flown = toFlow(
    "Transactionx",
    "sqlx",
    "*sqlx.Tx",
    wraps,
    `func (tx *Tx) BindNamed(query string, arg interface{}) (string, []interface{}, error)
func (tx *Tx) DriverName() string
func (tx *Tx) Get(dest interface{}, query string, args ...interface{}) error
func (tx *Tx) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
func (tx *Tx) MustExec(query string, args ...interface{}) sql.Result
func (tx *Tx) MustExecContext(ctx context.Context, query string, args ...interface{}) sql.Result
func (tx *Tx) NamedExec(query string, arg interface{}) (sql.Result, error)
func (tx *Tx) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
func (tx *Tx) NamedQuery(query string, arg interface{}) (*Rows, error)
func (tx *Tx) NamedStmt(stmt *NamedStmt) *NamedStmt
func (tx *Tx) NamedStmtContext(ctx context.Context, stmt *NamedStmt) *NamedStmt
func (tx *Tx) PrepareNamed(query string) (*NamedStmt, error)
func (tx *Tx) PrepareNamedContext(ctx context.Context, query string) (*NamedStmt, error)
func (tx *Tx) Preparex(query string) (*Stmt, error)
func (tx *Tx) PreparexContext(ctx context.Context, query string) (*Stmt, error)
func (tx *Tx) QueryRowx(query string, args ...interface{}) *Row
func (tx *Tx) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *Row
func (tx *Tx) Queryx(query string, args ...interface{}) (*Rows, error)
func (tx *Tx) QueryxContext(ctx context.Context, query string, args ...interface{}) (*Rows, error)
func (tx *Tx) Rebind(query string) string
func (tx *Tx) Select(dest interface{}, query string, args ...interface{}) error
func (tx *Tx) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
func (tx *Tx) Stmtx(stmt interface{}) *Stmt
func (tx *Tx) StmtxContext(ctx context.Context, stmt interface{}) *Stmt
func (tx *Tx) Unsafe() *Tx`
  );
  flown.functions.push(...tx().functions);
  return flown;
};

const stmt = () => {
  let flown = toFlow(
    "Stmt",
    "sql",
    "*sql.Stmt",
    wraps,
    `func (s *Stmt) Close() error
func (s *Stmt) Exec(args ...any) (Result, error)
func (s *Stmt) ExecContext(ctx context.Context, args ...any) (Result, error)
func (s *Stmt) Query(args ...any) (*Rows, error)
func (s *Stmt) QueryContext(ctx context.Context, args ...any) (*Rows, error)
func (s *Stmt) QueryRow(args ...any) *Row
func (s *Stmt) QueryRowContext(ctx context.Context, args ...any) *Row`
  );
  flown.functions.push(...tx().functions);
  return flown;
};

const stmtx = () => {
  let flown = toFlow(
    "Stmtx",
    "sqlx",
    "*sqlx.Stmt",
    wraps,
    `func (s *Stmt) Get(dest interface{}, args ...interface{}) error
func (s *Stmt) GetContext(ctx context.Context, dest interface{}, args ...interface{}) error
func (s *Stmt) MustExec(args ...interface{}) sql.Result
func (s *Stmt) MustExecContext(ctx context.Context, args ...interface{}) sql.Result
func (s *Stmt) QueryRowx(args ...interface{}) *Row
func (s *Stmt) QueryRowxContext(ctx context.Context, args ...interface{}) *Row
func (s *Stmt) Queryx(args ...interface{}) (*Rows, error)
func (s *Stmt) QueryxContext(ctx context.Context, args ...interface{}) (*Rows, error)
func (s *Stmt) Select(dest interface{}, args ...interface{}) error
func (s *Stmt) SelectContext(ctx context.Context, dest interface{}, args ...interface{}) error
func (s *Stmt) Unsafe() *Stmt`
  );
  flown.functions.push(...stmt().functions);
  return flown;
};

const result = () => {
  let flown = toFlow(
    "Result",
    "sql",
    "sql.Result",
    wraps,
    `func (s *Result) LastInsertId() (int64, error)
func (s *Result) RowsAffected() (int64, error)`
  );
  return flown;
};

console.log(JSON.stringify(result()));
