/**
 * OPERA Cloud Cashiering API
 * APIs to cater for Cashiering related external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Used when a message has been successfully processed to report any warnings or business errors that occurred.
 * @export
 * @interface WarningType
 */
export interface WarningType {
    /**
     * Property Value
     * @type {string}
     * @memberof WarningType
     */
    value?: string;
    /**
     * An abbreviated version of the error in textual format.
     * @type {string}
     * @memberof WarningType
     */
    shortText?: string;
    /**
     * If present, this refers to a table of coded values exchanged between applications to identify errors or warnings.
     * @type {string}
     * @memberof WarningType
     */
    code?: string;
    /**
     * If present, this URL refers to an online description of the error that occurred.
     * @type {string}
     * @memberof WarningType
     */
    docURL?: string;
    /**
     * If present, recommended values are those enumerated in the ErrorRS, (NotProcessed Incomplete Complete Unknown) however, the data type is designated as string data, recognizing that trading partners may identify additional status conditions not included in the enumeration.
     * @type {string}
     * @memberof WarningType
     */
    status?: string;
    /**
     * If present, this attribute may identify an unknown or misspelled tag that caused an error in processing. It is recommended that the Tag attribute use XPath notation to identify the location of a tag in the event that more than one tag of the same name is present in the document. Alternatively, the tag name alone can be used to identify missing data [Type=ReqFieldMissing].
     * @type {string}
     * @memberof WarningType
     */
    tag?: string;
    /**
     * If present, this attribute allows for batch processing and the identification of the record that failed amongst a group of records. This value may contain a concatenation of a unique failed transaction ID with specific record(s) associated with that transaction.
     * @type {string}
     * @memberof WarningType
     */
    recordId?: string;
    /**
     * The Warning element MUST contain the Type attribute that uses a recommended set of values to indicate the warning type. The validating XSD can expect to accept values that it has NOT been explicitly coded for and process them by using Type ="Unknown".
     * @type {string}
     * @memberof WarningType
     */
    type?: string;
    /**
     * Language identification.
     * @type {string}
     * @memberof WarningType
     */
    language?: string;
    /**
     * Reference Place Holder used as an index for this warning.
     * @type {string}
     * @memberof WarningType
     */
    rph?: string;
}
/**
 * Check if a given object implements the WarningType interface.
 */
export declare function instanceOfWarningType(value: object): boolean;
export declare function WarningTypeFromJSON(json: any): WarningType;
export declare function WarningTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): WarningType;
export declare function WarningTypeToJSON(value?: WarningType | null): any;
