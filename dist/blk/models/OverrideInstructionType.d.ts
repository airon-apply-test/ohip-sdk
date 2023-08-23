/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Type for Overrides. Contains information for the override action performed while booking a reservation.
 * @export
 * @interface OverrideInstructionType
 */
export interface OverrideInstructionType {
    /**
     * The description of the restriction for which the override was done.
     * @type {string}
     * @memberof OverrideInstructionType
     */
    description?: string;
    /**
     * The date when the override was done.
     * @type {Date}
     * @memberof OverrideInstructionType
     */
    date?: Date;
    /**
     * The type of override done. If done for Availability, then it will be AVAILABILITY.
     * @type {string}
     * @memberof OverrideInstructionType
     */
    type?: string;
    /**
     * Login ID of the user who performed the override.
     * @type {string}
     * @memberof OverrideInstructionType
     */
    userId?: string;
}
/**
 * Check if a given object implements the OverrideInstructionType interface.
 */
export declare function instanceOfOverrideInstructionType(value: object): boolean;
export declare function OverrideInstructionTypeFromJSON(json: any): OverrideInstructionType;
export declare function OverrideInstructionTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): OverrideInstructionType;
export declare function OverrideInstructionTypeToJSON(value?: OverrideInstructionType | null): any;
