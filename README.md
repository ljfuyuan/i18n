i18n
====

Package i18n is for app Internationalization and Localization.

## Usage

First of all, import this package:

```go
import "github.com/ljfuyuan/i18n"
```

The format of locale files is INI format configuration file, And file's name must be like this `zh_CN.ini`

## Minimal example

Here are two simplest locale file examples:

File `en-US.ini`:

```ini
hi = hello, %s
bye = goodbye
```

File `zh-CN.ini`:

```ini
hi = 您好，%s
bye = 再见
```

### Translation

```go
if err := i18n.Init("iniDirPath","en_Us");err !=nil {
    return err
}
i18n.Tr("en_US", "hi", "fuyuan")
i18n.Tr("en_US", "bye")
```

Code above will produce correspondingly:

- English `en-US`：`hello, fuyuan`, `goodbye`
- Chinese `zh-CN`：`您好，fuyuan`, `再见`

## Section

i18n module also uses the section feature of INI format configuration

For example

Content in locale file:

```ini
hi = 你好

[section]
hi = 你好,golang
```

Get `hi` in default:

```go
i18n.Tr("zh_CN", "hi")
```

Get `about` in about page:

```go
i18n.Tr("zh_CN", "section.hi")
```

## More information

- When matching non-default locale and didn't find the string, i18n will have a second try on default locale.
- If i18n still cannot find string in the default locale, raw string will be returned. For instance, when the string is `hi` and it does not exist in locale file, simply return `hi` as output.
