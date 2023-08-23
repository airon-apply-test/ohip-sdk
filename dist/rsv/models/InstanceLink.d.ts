/**
 * OPERA Cloud Reservation API
 * APIs to cater for Reservation functionality in OPERA Cloud. <br /><br />OPERA Cloud Reservations provides a complete set of capabilities for creating and updating reservations. Reservations are a central feature of OPERA Cloud. As a key source of information, the reservation specifies a guest&apos;s arrival date, departure date, room type, rate, packages, and many other details. It is also a gateway to dozens of other functions that contribute to the guest&apos;s experience.  All reservations in OPERA Cloud require a guest profile.<br /><br /> You can create profiles while booking a reservation. If a profile already exists, you can look it up (using getProfiles in the Customer Relationship Management module) and attach it to the reservation during the reservation booking process using the profileId.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Metadata describing link description objects that MAY appear in the JSON instance representation.
 * @export
 * @interface InstanceLink
 */
export interface InstanceLink {
    /**
     * URI [RFC3986] or URI Template [RFC6570]. If the value is set to URI Template, then the "templated" property must be set to true.
     * @type {string}
     * @memberof InstanceLink
     */
    href: string;
    /**
     * Name of the link relation that, in addition to the type property, can be used to retrieve link details. For example, href or profile.
     * @type {string}
     * @memberof InstanceLink
     */
    rel: string;
    /**
     * Boolean flag that specifies that "href" property is a URI or URI Template. If the property is a URI template, set this value to true. By default, this value is false.
     * @type {boolean}
     * @memberof InstanceLink
     */
    templated?: boolean;
    /**
     * HTTP method for requesting the target of the link.
     * @type {string}
     * @memberof InstanceLink
     */
    method: InstanceLinkMethodEnum;
    /**
     * Link to the metadata of the resource, such as JSON-schema, that describes the resource expected when dereferencing the target resource..
     * @type {string}
     * @memberof InstanceLink
     */
    targetSchema?: string;
    /**
     * The operationId of the path you can call to follow this link.  This allows you to look up not only the path and method, but the description of that path and any parameters you need to supply.
     * @type {string}
     * @memberof InstanceLink
     */
    operationId: string;
    /**
     * Exact copy of the "summary" field on the linked operation.
     * @type {string}
     * @memberof InstanceLink
     */
    title?: string;
}
/**
 * @export
 */
export declare const InstanceLinkMethodEnum: {
    readonly Get: "GET";
    readonly Post: "POST";
    readonly Put: "PUT";
    readonly Delete: "DELETE";
    readonly Patch: "PATCH";
    readonly Options: "OPTIONS";
    readonly Head: "HEAD";
};
export type InstanceLinkMethodEnum = typeof InstanceLinkMethodEnum[keyof typeof InstanceLinkMethodEnum];
/**
 * Check if a given object implements the InstanceLink interface.
 */
export declare function instanceOfInstanceLink(value: object): boolean;
export declare function InstanceLinkFromJSON(json: any): InstanceLink;
export declare function InstanceLinkFromJSONTyped(json: any, ignoreDiscriminator: boolean): InstanceLink;
export declare function InstanceLinkToJSON(value?: InstanceLink | null): any;
