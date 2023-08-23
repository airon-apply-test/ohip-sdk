/**
 * OPERA Cloud CRM Configuration API
 * APIs for Customer Relationship Management (profile) configuration, such as creating preferences, or address types.  It also includes Membership Configuration, where you can retrieve membership levels that are configured for a property, or create new membership enrollment codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Applies on restricted dates.
 * @export
 */
export declare const RuleAppliesOnType: {
    readonly Arrival: "Arrival";
    readonly Departure: "Departure";
    readonly Any: "Any";
    readonly Restricted: "Restricted";
};
export type RuleAppliesOnType = typeof RuleAppliesOnType[keyof typeof RuleAppliesOnType];
export declare function RuleAppliesOnTypeFromJSON(json: any): RuleAppliesOnType;
export declare function RuleAppliesOnTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RuleAppliesOnType;
export declare function RuleAppliesOnTypeToJSON(value?: RuleAppliesOnType | null): any;
