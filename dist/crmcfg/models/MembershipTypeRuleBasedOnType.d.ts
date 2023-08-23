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
 * Indicates Membership upgrades/downgrades to the next tier level uses RFM (Recency, Frequency, Monetary Value) score.
 * @export
 */
export declare const MembershipTypeRuleBasedOnType: {
    readonly Revenue: "Revenue";
    readonly Stay: "Stay";
    readonly Nights: "Nights";
    readonly Enrollment: "Enrollment";
    readonly TierUpgrade: "TierUpgrade";
    readonly Renewal: "Renewal";
    readonly Rfm: "Rfm";
};
export type MembershipTypeRuleBasedOnType = typeof MembershipTypeRuleBasedOnType[keyof typeof MembershipTypeRuleBasedOnType];
export declare function MembershipTypeRuleBasedOnTypeFromJSON(json: any): MembershipTypeRuleBasedOnType;
export declare function MembershipTypeRuleBasedOnTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): MembershipTypeRuleBasedOnType;
export declare function MembershipTypeRuleBasedOnTypeToJSON(value?: MembershipTypeRuleBasedOnType | null): any;
