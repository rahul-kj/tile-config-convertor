Tile Configuration Convertor
---

## Motivation
To automate deployment of any of the product tiles shipped by Pivotal, the configuration parameters are the ones that are hard to fetch. This tool was written to ease the fetching of the properties of the product tile, after the tile has been uploaded and staged.

## How this works
Once the product is uploaded and staged, the platform engineer can use the [om cli](https://github.com/pivotal-cf/om/releases).

To download all the available properties for a given tile execute:

```
om -t OPS-MANAGER-URL -u USERNAME -p PASSWORD curl -p /api/v0/staged/products/$PRODUCT_GUID/properties
```

To download all the available resources for a given tile execute:

```
om -t OPS-MANAGER-URL -u USERNAME -p PASSWORD curl -p /api/v0/staged/products/$PRODUCT_GUID/resources
```

These commands produce a **json** file.

In-order to get the right set of properties that can be configured, execute

```
tile-config-convertor -g properties -i properties.json -o properties.yml
```

In-order to get the right set of properties that can be configured, execute

```
tile-config-convertor -g resources -i resources.json -o resources.yml
```

You can now paste the output contents into the params.yml for the given tile and fly them using the [install-product pipeline](https://github.com/rahul-kj/pcf-concourse-pipelines/tree/master/pipelines/install-product)

**NOTE: New features will be added as and when they are identified**
